package gitapiclient

import (
	"strings"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
	configuration "xqledger/apirouter/configuration"
	utils "xqledger/apirouter/utils"

	resty "gopkg.in/resty.v1"
)

const componentMessage = "Gitea API Client"
const okcode = 200
const createdcode = 201
var config = configuration.GlobalConfiguration
var client *resty.Client

func getAPIClient() *resty.Client{
	if client != nil {
		return client
	}
	client = resty.New()
	client.SetTimeout(time.Duration(config.Gitserver.Strategy.Timeout) * time.Millisecond)
	client.SetHeaders(map[string]string{
        "Content-Type": "application/json",
        "User-Agent": "APIRouter",
		"Authorization": config.Gitserver.Authtoken,
    })
	return client
}

func StartSession(repoName, branchName string) (string, error) {
	methodMessage := "StartSession"
	sessionBranchName, sessionBranchErr := createSessionBranch(repoName, branchName)
	if sessionBranchErr != nil {
		utils.PrintLogError(sessionBranchErr, componentMessage, methodMessage, "Error creating new branch")
		return "", sessionBranchErr
	}
	return sessionBranchName, nil
}

func getNameFromEmail(email string) (string, error) {
	at := strings.LastIndex(email, "@")
    if at >= 0 {
        username, _ := email[:at], email[at+1:]
		return username, nil
    } else {
		err := errors.New(fmt.Sprintf("Error: %s is an invalid email address\n", email))
		return "", err
	}
}

func CommitSession(repoName string, session utils.Session) error{
	methodMessage := "CommitSession"
	now := time.Now()
	var prRequest CreatePullRequestOption
	prRequest.Head = session.Branch
	prRequest.Base = "master"
	prRequest.Title = strconv.FormatInt(session.SessionID, 10)
	user, userErr := getNameFromEmail(session.User)
	if userErr != nil {
		utils.PrintLogError(userErr, componentMessage, methodMessage, "Error getting username for the new pull request")
		return userErr
	}
	prRequest.Assignee = user
	prRequest.Body = session.Description
	prRequest.Due_date = now.Format(time.RFC3339)
	prRequest.Milestone = 0

	// index, prErr := createPullRequest(repoName, prRequest)
	// if prErr != nil {
	// 	utils.PrintLogError(prErr, componentMessage, methodMessage, "Error creating new pull request")
	// 	return prErr
	// }

	prCommit, prCommirErr := getLastCommitInPR(repoName, 6)
	if prCommirErr != nil {
		utils.PrintLogError(prCommirErr, componentMessage, methodMessage, "Error getting last commit in PR")
		return prCommirErr
	}
	fmt.Println("COMMIT")
	fmt.Println(prCommit)
	var mergeRequestOption MergePullRequestOption
	mergeRequestOption.Do = "merge" // 	[ merge, rebase, rebase-merge, squash, manually-merged ]
	mergeRequestOption.MergeMessageField = session.Description
	mergeRequestOption.MergeTitleField = "Session ID " + strconv.FormatInt(session.SessionID, 10)
	mergeRequestOption.Delete_branch_after_merge = true
	mergeRequestOption.Force_merge = true
	mergeRequestOption.Head_commit_id = prCommit
	mergeRequestOption.MergeCommitID = prCommit

	mergeErr := mergePullRequest(repoName, 6, mergeRequestOption)
	if mergeErr != nil {
		utils.PrintLogError(mergeErr, componentMessage, methodMessage, "Error creating new merge request")
		return mergeErr
	}
	return nil
}

func getLastCommitInPR(repoName string, index int64) (string, error) {
	c := getAPIClient()
	resp, err := c.R().
		Get(fmt.Sprintf("%s/repos/%s/%s/pulls/%d/commits", config.Gitserver.Url, config.Gitserver.Username, repoName, index))

	if err != nil {
		return "", err
	}
	if resp.StatusCode() != okcode {
		err := errors.New(fmt.Sprintf("Unexpected status code, expected %d, got %d instead", createdcode, resp.StatusCode()))
		return "", err
	}
	commits := []PullRequestCommit{}
	unmarshalErr := json.Unmarshal(resp.Body(), &commits)
	if unmarshalErr != nil {
		fmt.Println(unmarshalErr)
		return "", unmarshalErr
	}
	sha := commits[0].Sha
	return sha, nil
}

func createSessionBranch(repoName, branchName string) (string, error) {
	var branchRequest BranchRequest
	branchRequest.New_branch_name = branchName
	branchRequest.Old_branch_name = ""
	c := getAPIClient()
	resp, err := c.R().
		SetBody(branchRequest).
		Post(fmt.Sprintf("%s/repos/%s/%s/branches", config.Gitserver.Url, config.Gitserver.Username, repoName))

	if err != nil {
		return "", err
	}
	if resp.StatusCode() != createdcode {
		err := errors.New(fmt.Sprintf("Unexpected status code, expected %d, got %d instead", createdcode, resp.StatusCode()))
		return "", err
	}
	return branchName, nil
}

func createPullRequest(repoName string, pr CreatePullRequestOption) (int64, error){
	c := getAPIClient()
	var result PullRequest

	resp, err := c.R().
		SetBody(pr).
		Post(fmt.Sprintf("%s/repos/%s/%s/pulls", config.Gitserver.Url, config.Gitserver.Username, repoName))

	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	if resp.StatusCode() != createdcode {
		err := errors.New(fmt.Sprintf("Unexpected status code, expected %d, got %d instead", createdcode, resp.StatusCode()))
		return 0, err
	}
	
	unmarshalErr := json.Unmarshal(resp.Body(), &result)

	if unmarshalErr != nil {
		fmt.Println(unmarshalErr.Error())
		return 0,unmarshalErr
	}
	return result.Id, nil
}

func mergePullRequest(repoName string, prIndex int64, mpr MergePullRequestOption) error {
	mpr.Delete_branch_after_merge = config.Gitserver.Strategy.Deletebranchaftermerge
	mpr.Force_merge = config.Gitserver.Strategy.Forcemerge
	c := getAPIClient()
	fmt.Println(fmt.Sprintf("%s/repos/%s/%s/pulls/%d/merge", config.Gitserver.Url, config.Gitserver.Username, repoName, prIndex))
	_, err := c.R().
		SetBody(mpr).
		Post(fmt.Sprintf("%s/repos/%s/%s/pulls/%d/merge", config.Gitserver.Url, config.Gitserver.Username, repoName, prIndex))
	if err != nil {
		fmt.Println("ERROR MERGE: " + err.Error())
		return err
	}
	return nil
}

type CreatePullRequestOption struct {
	Assignee   string `json:"assignee"`
	Assignees   []string `json:"assignees"`
	Base   string `json:"base"`
	Body   string `json:"body"`
	Due_date   string `json:"due_date"`
	Head   string `json:"head"`
	Labels   []int64 `json:"labels"`
	Milestone   int64 `json:"milestone"`
	Title   string `json:"title"`
}

type Identity struct {
	Email   string `json:"email"`
	Name   string `json:"name"`
}

type PullRequest struct {
	Id   int64 `json:"id"`
}

type Session struct {
	Id   string `json:"id"`
	Starting_time   string `json:"starting_time"`
}

type PullRequestCommit struct {
	Sha   string `json:"sha"`
}

type BranchRequest struct {
	New_branch_name   string `json:"new_branch_name"`
	Old_branch_name   string `json:"old_branch_name"`
}

type BranchResponse struct {
	Effective_branch_protection_name   string `json:"effective_branch_protection_name"`
	name   string `json:"name"`
}

type MergePullRequestOption struct {
	Do   string `json:"do"` // [ merge, rebase, rebase-merge, squash, manually-merged ]
	MergeCommitID   string `json:"MergeCommitID"`
	MergeMessageField string `json:"MergeMessageField"`
	MergeTitleField string `json:"MergeTitleField"`
	Delete_branch_after_merge bool `json:"delete_branch_after_merge"`
	Force_merge bool `json:"force_merge"`
	Head_commit_id string `json:"head_commit_id"`
}


type CreateRepoOption struct {
	Auto_init   bool `json:"auto_init"` // true
	Default_branch   string `json:"default_branch"` // master
	Description   string `json:"description"`
	Gitignores   string `json:"gitignores"` // gitignores to use
	Issue_labels   string `json:"issue_labels"` // Label-Set to use
	License   string `json:"license"`
	Name   string `json:"name"` // must be UNIQUE in server
	Private   bool `json:"name"` // true
	Readme   string `json:"readme"` // Readme of the repository to create
	Template   bool `json:"template"` // false
	Trust_model   string `json:"trust_model"` // default, collaborator, committer, collaboratorcommitter
}


func CreateNewRepo (name, description string) error {
	var repo CreateRepoOption
	repo.Name = name
	repo.Description = "New Test Repo"
	repo.Issue_labels = ""
	repo.Auto_init = true
	repo.Default_branch = "master"
	repo.Gitignores = ""
	repo.License = ""
	repo.Private = true
	repo.Readme = ""
	repo.Template = false
	repo.Trust_model = "default"

	c := getAPIClient()

	resp, err := c.R().
		SetBody(repo).
		Post(fmt.Sprintf("%s/user/repos/", config.Gitserver.Url))

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	if resp.StatusCode() != createdcode {
		err := errors.New(fmt.Sprintf("Unexpected status code, expected %d, got %d instead", createdcode, resp.StatusCode()))
		return err
	}
	return nil
}