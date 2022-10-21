package admindb

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	utils "xqledger/apirouter/utils"
)

const dbowner = "TestOrchestrator"
const dbname = "TestRepository"
const collection = "main"
const collectionID = 1
const repo = "GitOperatorTestRepo"
const id = "012345678912345678912345"
const commit = "2077eb80abf9522a163c7ce012c8b90dd87d5a8c"
const email = "testorchestrator@gmail.com"
const recordTime = int64(1636570869)
var sessionID = utils.GetRandomIDInt()
var branchName = utils.GetUUID()

func TestGetTenants(t *testing.T) {
	Convey("Check GetTenants ", t, func() {
		tenants, err := GetTenants()
		So(err, ShouldBeNil)
		So(len(tenants), ShouldBeGreaterThan, 0)
	})
}

func TestAllOpsWithSessionsInDB(t *testing.T) {
	var candidateSession utils.Session
	var col utils.Collection
	col.Active = true
	col.CollectionID = 1
	col.Creation = 1641035236
	col.Description = "First collection for testing"
	col.Name = collection
	candidateSession.Collection = col
	candidateSession.Description = "Description for a test session"
	candidateSession.Branch = branchName
	candidateSession.EndTime = 0
	candidateSession.StartTime = 1641748556
	candidateSession.User = email
	candidateSession.SessionID = sessionID

	Convey("Check Open a session ", t, func() {
		newsession, err := NewSession(candidateSession)
		So(err, ShouldBeNil)
		So(newsession.SessionID, ShouldBeGreaterThan, 0)
	})

	Convey("Check get session for a collection by ID", t, func() {
		session, err := GetSessionByID(16)
		So(err, ShouldBeNil)
		So(session.SessionID, ShouldEqual, 16)
	})

	Convey("Check get all sessions for a collection ", t, func() {
		sessions, err := GetAllSessionsByCollection(collectionID)
		So(err, ShouldBeNil)
		So(len(sessions), ShouldBeGreaterThan, 0)
	})

	Convey("Check get acxtive sessions for a collection ", t, func() {
		sessions, err := GetActiveSessionsByCollection(collectionID)
		So(err, ShouldBeNil)
		So(len(sessions), ShouldBeGreaterThan, 0)
	})

	Convey("Check close session ", t, func() {
		err := CloseSession(sessionID)
		So(err, ShouldBeNil)
	})
}
