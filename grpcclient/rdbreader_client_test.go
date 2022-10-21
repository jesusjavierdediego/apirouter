package grpcclient

import (
	// "os"
	// "testing"
	// . "github.com/smartystreets/goconvey/convey"
)

// func TestGetHIstory(t *testing.T) {
// 	Convey("Getting HIstory of a record ", t, func() {
// 		os.Setenv("PROFILE", "dev")
// 		history, err := GetRecordHistory("2.json", "gitrepo")
// 		So(err, ShouldBeNil)
// 		So(len(history.Commits), ShouldBeGreaterThan, 0)
// 	})
// }

// func TestGetContentInCommit(t *testing.T) {
// 	Convey("Getting content in a commit of a record ", t, func() {
// 		os.Setenv("PROFILE", "dev")
// 		err := GetContentInCommit("2077eb80abf9522a163c7ce012c8b90dd87d5a8c", "2.json", "gitrepo")
// 		So(err, ShouldBeNil)
// 	})
// }

// func TestGetDiff(t *testing.T) {
// 	Convey("Getting diff between two commits of a record ", t, func() {
// 		os.Setenv("PROFILE", "dev")
// 		err := GetDiffTwoCommitsInFile("30bf07336744d15b88cec5197fb0bd05991a6dfd", "2077eb80abf9522a163c7ce012c8b90dd87d5a8c", "2.json", "gitrepo")
// 		So(err, ShouldBeNil)
// 	})
// }