package grpcclient

import (
	"fmt"
	"context"
	"log"
	"time"
	"strconv"
    "google.golang.org/grpc"
	pb "xqledger/apirouter/protobuf"
	configuration "xqledger/apirouter/configuration"
	utils "xqledger/apirouter/utils"
)

var config = configuration.GlobalConfiguration
var gitreader_address = config.Grpcclient.Gitreaderhost + ":" + strconv.Itoa(config.Grpcclient.Gitreaderport)
var gitreader_conn *grpc.ClientConn
var gitreader_connErr error
const componentMessage = "GRPC Client"

func getGitReaderServerConn() (*grpc.ClientConn, error){
	gitreader_conn, gitreader_connErr = grpc.Dial(gitreader_address, grpc.WithInsecure())
	if gitreader_connErr != nil {
		log.Fatalf("did not connect: %v", gitreader_connErr)
		return nil, gitreader_connErr
	}
	return gitreader_conn, nil
}

func GetRecordHistory(filepath, reponame string) (*pb.RecordHistory, error){
	var methodMessage = "GetRecordHistory"
	var result pb.RecordHistory
	gitreader_conn, gitreader_connErr = getGitReaderServerConn()
	if gitreader_connErr != nil {
		utils.PrintLogError(gitreader_connErr, componentMessage, methodMessage, "Error in connection")
		return &result, gitreader_connErr
	}
	defer gitreader_conn.Close()
	c := pb.NewRecordHistoryServiceClient(gitreader_conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var q pb.Query
	q.CommitIdOld = ""
	q.CommitIdNew = ""
	q.FilePath = filepath
	q.RepoName = reponame

	history, err := c.GetRecordHistory(ctx, &q)
	if err != nil {
		utils.PrintLogError(gitreader_connErr, componentMessage, methodMessage, fmt.Sprintf("could not get history - Reason: %v", err))
		return &result, err
	}
	utils.PrintLogInfo(componentMessage, methodMessage, fmt.Sprintf("SUCCESS!! History: %s", history.Commits))
	return history, nil
}

func GetContentInCommit(commit_id, filepath, reponame string) (*pb.CommitContent, error){
	var methodMessage = "GetContentInCommit"
	var result pb.CommitContent
	gitreader_conn, gitreader_connErr = getGitReaderServerConn()
	if gitreader_connErr != nil {
		utils.PrintLogError(gitreader_connErr, componentMessage, methodMessage, "Error in connection")
		return &result, gitreader_connErr
	}
	defer gitreader_conn.Close()
	c := pb.NewRecordHistoryServiceClient(gitreader_conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var q pb.Query
	q.CommitIdOld = commit_id
	q.CommitIdNew = ""
	q.FilePath = filepath
	q.RepoName = reponame

	content, err := c.GetContentInCommit(ctx, &q)
	if err != nil {
		utils.PrintLogError(gitreader_connErr, componentMessage, methodMessage, fmt.Sprintf("could not get content - Reason: %v", err))
		return &result, err
	}
	utils.PrintLogInfo(componentMessage, methodMessage, fmt.Sprintf("SUCCESS!! Content: %s", content.Content))
	return content, nil
}

func GetDiffTwoCommitsInFile(commit_id_old, commit_id_new, filepath, reponame string) (*pb.DiffHtml, error){
	var methodMessage = "GetDiffTwoCommitsInFile"
	var result pb.DiffHtml
	gitreader_conn, gitreader_connErr = getGitReaderServerConn()
	if gitreader_connErr != nil {
		utils.PrintLogError(gitreader_connErr, componentMessage, methodMessage, "Error in connection")
		return &result, gitreader_connErr
	}
	defer gitreader_conn.Close()
	c := pb.NewRecordHistoryServiceClient(gitreader_conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var q pb.Query
	q.CommitIdOld = commit_id_old
	q.CommitIdNew = commit_id_new
	q.FilePath = filepath
	q.RepoName = reponame

	diff, err := c.GetDiffTwoCommitsInFile(ctx, &q)
	if err != nil {
		utils.PrintLogError(gitreader_connErr, componentMessage, methodMessage, fmt.Sprintf("could not get diff - Reason: %v", err))
		return &result, err
	}
	utils.PrintLogInfo(componentMessage, methodMessage, fmt.Sprintf("SUCCESS!! Diff: %s", diff.Html))
	return diff, nil
}