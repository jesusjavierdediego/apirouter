package gitapiclient

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	utils "xqledger/apirouter/utils"
)

/*
func StartSession(repoName, branchName string) (string, error) {
func CommitSession(repoName string, session utils.Session) error{


*/

const dbowner = "TestOrchestrator"
const dbname = "TestRepository"
const collection = "main"
const collectionID = 1
const repo = "GitOperatorTestRepo"
const id = "012345678912345678912345"
const commit = "2077eb80abf9522a163c7ce012c8b90dd87d5a8c"
const email = "testorchestrator@gmail.com"
const recordTime = int64(1636570869)
const sessionID = 1

func TestSessions(t *testing.T) {
	var branchName = utils.GetUUID()
	var candidateSession utils.Session
	var col utils.Collection
	col.Active = true
	col.CollectionID = 1
	col.Creation = 0
	col.Description = "Test collection"
	col.Name = collection
	candidateSession.Collection = col
	candidateSession.Description = ""
	candidateSession.Branch = branchName
	candidateSession.EndTime = 0
	candidateSession.StartTime = 0
	candidateSession.User = email
	candidateSession.SessionID = sessionID
	Convey("Check StartSession ", t, func() {
		bName, err := StartSession(dbname, branchName)
		So(err, ShouldBeNil)
		So(bName, ShouldEqual, branchName)
	})
	Convey("Check CommitSession ", t, func() {
		err := CommitSession(dbname, candidateSession)
		So(err, ShouldBeNil)
	})
}