package rest

import (
	"bytes"
	"fmt"
	"net/http"
    "net/http/httptest"
	"testing"
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/gin-gonic/gin"
	utils "xqledger/apirouter/utils"
)

const dbowner = "TestOrchestrator"
const dbname = "TestRepository"
const collection = "main"
const repo = "GitOperatorTestRepo"
const id = "012345678912345678912345"
const email = "testorchestrator@gmail.com"
const recordTime = int64(1636570869)
const file = "17155780968166547879.json"
const commitnew = "ce6f274afa3549be010fd8c3d25d5b2d10ba56e4"
const commitold = "15f0f595672f500aad6648b2b522020705de8147"


func getEvent()utils.RecordEvent{
	record := utils.RecordEvent{}
	record.Id = id
	record.Group = ""
	record.DBName = repo
	record.User = email
	record.OperationType = "new"
	record.SendingTime = recordTime
	record.ReceptionTime = recordTime
	record.ProcessingTime = recordTime
	record.Priority = "MEDIUM"
	record.RecordContent = "{\"browsers\":{\"firefox\":{\"name\":\"Firefox\",\"pref_url\":\"about:config\",\"releases\":{\"1\":{\"release_date\":\"2004-11-09\",\"status\":\"retired\",\"engine\":\"Gecko\",\"engine_version\":\"1.7\"}}}}}"
	record.Status = "PENDING"
	return record
}

func getUpdateEvent()utils.RecordEvent{
	record := utils.RecordEvent{}
	record.Id = id
	record.Group = ""
	record.DBName = repo
	record.User = email
	record.OperationType = "update"
	record.SendingTime = recordTime
	record.ReceptionTime = recordTime
	record.ProcessingTime = recordTime
	record.Priority = "MEDIUM"
	record.RecordContent = "{\"browsers\":{\"firefox\":{\"name\":\"Firefox\",\"pref_url\":\"about:config\",\"releases\":{\"1\":{\"release_date\":\"2004-12-23\",\"status\":\"retired\",\"engine\":\"Gecko\",\"engine_version\":\"1.8\"}}}}}"
	record.Status = "PENDING"
	return record
}

func getEventBatch() utils.RecordEventBatch{
	var result utils.RecordEventBatch
	var records []utils.RecordEvent
	result.Id = "1"
	result.DBName = repo
	result.OperationType = "new"
	records = append(records, getEvent())
	records = append(records, getEvent())
	result.Records = records
	return result
}

func getUpdateEventBatch() utils.RecordEventBatch{
	var result utils.RecordEventBatch
	var records []utils.RecordEvent
	result.Id = "1"
	result.DBName = repo
	result.OperationType = "update"
	records = append(records, getUpdateEvent())
	records = append(records, getUpdateEvent())
	result.Records = records
	return result
}


func TestKeepAlive(t *testing.T) {
	Convey("Check keep alive ", t, func() {
		gin.SetMode(gin.TestMode)
		ts := httptest.NewServer(setupRouter())
		defer ts.Close()
		resp, _ := http.Get(fmt.Sprintf("%s/xqledger/v1/keepalive", ts.URL))
		So(resp.StatusCode, ShouldEqual, 200)
	})
}

// func getClient() *http.Client{
// 	client := &http.Client{
//         Timeout: time.Second * 10,
//     }
// 	return client
// }

func TestGetCountRecordsFromColl(t *testing.T) {
	Convey("Check GetCountInCollection ", t, func() {
		gin.SetMode(gin.TestMode)
		ts := httptest.NewServer(setupRouter())
		defer ts.Close()
		resp, _ := http.Get(fmt.Sprintf("%s/xqledger/v1/recordcount?dbname=%s&collection=%s", ts.URL, dbname, collection))
		So(resp.StatusCode, ShouldEqual, 200)
	})
}

// func TestGetRecordsByQuery(t *testing.T) {
// 	Convey("Check GetRecordsByQuery ", t, func() {
// 		gin.SetMode(gin.TestMode)
// 		ts := httptest.NewServer(setupRouter())
// 		defer ts.Close()

// 		var criteriaSet utils.CriteriaSet
// 		var criteria1 utils.Criteria
// 		criteria1.Booleanoperator = "AND"
// 		criteria1.Operator = ""
// 		criteria1.Parameter = ""
// 		criteria1.Value = ""
// 		criteriaSet.Set = append(criteriaSet.Set, criteria1)
// 		c, err := json.Marshal(criteriaSet)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		resp, _ := http.Post(fmt.Sprintf("%s/xqledger/v1/query/%s", ts.URL, dbname), "application/json",  bytes.NewBuffer(c))
// 		So(resp.StatusCode, ShouldEqual, 400)
// 	})
// }

func TestGetRecordHistory(t *testing.T) {
	Convey("Check GetRecordHistory ", t, func() {
		gin.SetMode(gin.TestMode)
		ts := httptest.NewServer(setupRouter())
		defer ts.Close()
		resp, _ := http.Get(fmt.Sprintf("%s/xqledger/v1/recordhistory", ts.URL))
		So(resp.StatusCode, ShouldEqual, 200)
	})
}

func TestGetContentInCommit(t *testing.T) {
	Convey("Check GetContentInCommit ", t, func() {
		gin.SetMode(gin.TestMode)
		ts := httptest.NewServer(setupRouter())
		defer ts.Close()
		var path = fmt.Sprintf("%s/xqledger/v1/recordevent?commit=%s&file=%s&db=%s", ts.URL,commitold,file,dbname)
		fmt.Println(path)
		resp, _ := http.Get(path)
		So(resp.StatusCode, ShouldEqual, 200)
	})
}

func TestGetDiffTwoCommitsInFile(t *testing.T) {
	Convey("Check GetDiffTwoCommitsInFile ", t, func() {
		gin.SetMode(gin.TestMode)
		ts := httptest.NewServer(setupRouter())
		defer ts.Close()
		resp, _ := http.Get(fmt.Sprintf("%s/xqledger/v1/recorddiff?commitnew=%s&file=%s&db=%s&commitold=%s", ts.URL,commitnew,file,dbname,commitold))
		So(resp.StatusCode, ShouldEqual, 200)
	})
}

func TestHandlePostPutRecord_1(t *testing.T) {
	Convey("Check HandlePostPutRecord ", t, func() {
		gin.SetMode(gin.TestMode)
		ts := httptest.NewServer(setupRouter())
		defer ts.Close()
		httpposturl := fmt.Sprintf("%s/xqledger/v1/record", ts.URL)
		fmt.Println("HTTP JSON POST URL:", httpposturl)
		event := getEvent()
		eventstring, err := json.Marshal(event)
		if err != nil {
			panic (err)
		}
		var jsonData = []byte(eventstring)
		request, error := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")

		client := &http.Client{}
		response, error := client.Do(request)
		if error != nil {
			panic(error)
		}
		defer response.Body.Close()
		So(response.StatusCode, ShouldEqual, 202)
	})
}

func TestHandlePostPutRecord_2(t *testing.T) {
	Convey("Check HandlePostPutRecord ", t, func() {
		gin.SetMode(gin.TestMode)
		ts := httptest.NewServer(setupRouter())
		defer ts.Close()
		httpposturl := fmt.Sprintf("%s/xqledger/v1/record", ts.URL)
		fmt.Println("HTTP JSON POST URL:", httpposturl)
		event := getUpdateEvent()
		eventstring, err := json.Marshal(event)
		if err != nil {
			panic (err)
		}
		var jsonData = []byte(eventstring)
		request, error := http.NewRequest("PUT", httpposturl, bytes.NewBuffer(jsonData))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")

		client := &http.Client{}
		response, error := client.Do(request)
		if error != nil {
			panic(error)
		}
		defer response.Body.Close()
		So(response.StatusCode, ShouldEqual, 202)
	})
}

func TestHandlePostPutBatch_1(t *testing.T) {
	Convey("Check HandlePostPutBatch Post ", t, func() {
		gin.SetMode(gin.TestMode)
		ts := httptest.NewServer(setupRouter())
		defer ts.Close()
		httpposturl := fmt.Sprintf("%s/xqledger/v1/batch", ts.URL)
		fmt.Println("HTTP JSON POST URL:", httpposturl)
		batch := getEventBatch()
		eventstring, err := json.Marshal(batch)
		if err != nil {
			panic (err)
		}
		var jsonData = []byte(eventstring)
		request, error := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")

		client := &http.Client{}
		response, error := client.Do(request)
		if error != nil {
			panic(error)
		}
		defer response.Body.Close()
		So(response.StatusCode, ShouldEqual, 202)
	})
}

func TestHandlePostPutBatch_2(t *testing.T) {
	Convey("Check HandlePostPutBatch Put ", t, func() {
		gin.SetMode(gin.TestMode)
		ts := httptest.NewServer(setupRouter())
		defer ts.Close()
		httpposturl := fmt.Sprintf("%s/xqledger/v1/batch", ts.URL)
		fmt.Println("HTTP JSON POST URL:", httpposturl)
		batch := getUpdateEventBatch()
		eventstring, err := json.Marshal(batch)
		if err != nil {
			panic (err)
		}
		var jsonData = []byte(eventstring)
		request, error := http.NewRequest("PUT", httpposturl, bytes.NewBuffer(jsonData))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")

		client := &http.Client{}
		response, error := client.Do(request)
		if error != nil {
			panic(error)
		}
		defer response.Body.Close()
		So(response.StatusCode, ShouldEqual, 202)
	})
}