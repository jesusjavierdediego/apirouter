package rest

import (
	"strconv"
	"encoding/json"
	"errors"
	"fmt"
	//"io/ioutil"
	kafka "xqledger/apirouter/kafka"
	utils "xqledger/apirouter/utils"
	gitapiclient "xqledger/apirouter/gitapiclient"
	grpcclient "xqledger/apirouter/grpcclient"
	admindb "xqledger/apirouter/admindb"
	"github.com/gin-gonic/gin"
)

const componentMessage = "REST Server Impl"

func KeepAlive(c *gin.Context) {
	c.JSON(200, "ok")
}

/*
Admin Section
*/
type owner struct {
    Repos []string
}


func getTenants(c *gin.Context){
	methodMessage := "getTenants"
	tenants, err := admindb.GetTenants()
	if err != nil {
		utils.PrintLogError(err, componentMessage, methodMessage, "Error retrieving Tenants")
		c.JSON(503, err.Error())
		return
	}
	utils.PrintLogInfo(componentMessage, methodMessage, "Tenants retrieved OK")
	c.JSON(200, tenants)
	return
}

func validateApiKey(c *gin.Context) {
	apikey := c.GetHeader("api-key")
	if !(len(apikey) >0) {
		msg := "API Key is empty"
		err := errors.New(msg)
		utils.PrintLogError(err, componentMessage, "validation of API Key", msg)
		c.JSON(400, msg)
		return
	}
}

func CreateNewDatabase(c *gin.Context) {
	/*
	1-Create new repo in git
	2-Create nbew collection in mongo
	3-Register db in admin db
	*/

	methodMessage := "CreateNewDatabase"
	// 1 create git repo
	name := c.Query("name")
	apikey := c.GetHeader("api-key")
	validateApiKey(c)
	if !(len(name) >0) {
		msg := "DB name is empty"
		err := errors.New(msg)
		utils.PrintLogError(err, componentMessage, methodMessage, msg)
		c.JSON(400, msg)
		return
	}
	description := "Standard description for new repo"
	gitErr := gitapiclient.CreateNewRepo(name, description)
	if gitErr != nil {
		utils.PrintLogError(gitErr, componentMessage, methodMessage, "Error creating DB")
		c.JSON(503, gitErr.Error())
		return
	}

	// 2 Create RDB collection
	// tODO

	// 3 Register new db
	tenants, err := admindb.GetTenants()
	if err != nil {
		utils.PrintLogError(err, componentMessage, methodMessage, "Error getting list of tenants")
		c.JSON(503, err.Error())
		return
	}
	var tenant utils.Tenant
	var tenantFound = false
	for _, t := range tenants {
		if strconv.Itoa(t.TenantID) == apikey {
			tenant = t
			tenantFound = true
		}
	}
	if !tenantFound {
		msg := "API Key is not valid (no match with list of tenants)"
		err := errors.New(msg)
		utils.PrintLogError(err, componentMessage, methodMessage, msg)
		c.JSON(403, msg)
		return
	}
	newdbErr := admindb.NewDB(name, description, tenant)
	if newdbErr != nil {
		utils.PrintLogError(newdbErr, componentMessage, methodMessage, "Error recording new db")
		c.JSON(503, newdbErr.Error())
		return
	}
	
	msg := "New DB created"
	utils.PrintLogInfo(componentMessage, methodMessage, msg)
	c.JSON(201, "New DB created")
}

func GetOpenSessions(c *gin.Context) {
	// TODO
	methodMsg := "GetOpenSessions"
	utils.PrintLogInfo(componentMessage, methodMsg, "")
	c.JSON(501, "Try later")
}

func GetAdminInfo(c *gin.Context) {
	// Composes the query and sends the query to the grpc server (RDBOperator)
	// TODO
	methodMsg := "GetAdminInfo"
	utils.PrintLogInfo(componentMessage, methodMsg, "")
	c.JSON(501, "Try later")
}

func GetAllRecords(c *gin.Context) {
	// Composes the query and sends the query to the grpc server (RDBOperator)
	// TODO
	methodMsg := "GetAllRecords"
	utils.PrintLogInfo(componentMessage, methodMsg, utils.Query_byID_ok)
	c.JSON(501, "Try later")
}

func GetRecordById(c *gin.Context) {
	// Composes the query and sends the query to the grpc server (RDBOperator)
	// TODO
	methodMsg := "GetRecordById"
	utils.PrintLogInfo(componentMessage, methodMsg, utils.Query_byID_ok)
	c.JSON(501, "Try later")
}

func DeleteRecordById(c *gin.Context) {
	// TODO
	methodMsg := "DeleteRecordById"
	utils.PrintLogInfo(componentMessage, methodMsg, utils.Event_delete_ok)	
	c.JSON(501, "Try later")
}

/*
End Admin Section
*/



type CountCollection struct {
	Total   int64 `json:"count"`
}

func GetCountInCollection(c *gin.Context) {
	dbname := c.Query("dbname")
	collection := c.Query("collection")
	if !(len(dbname)>0) || !(len(collection)> 0) {
		c.JSON(400, "Query parameters are not right")
		return
	}
	count, err := grpcclient.GetCountFromColl(dbname, collection)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	var result CountCollection
	result.Total = count
	c.JSON(200, result)
	return
}

func GetRecordHistory(c *gin.Context) {
	filepath := c.Query("file")
	reponame := c.Query("db")

	history, err := grpcclient.GetRecordHistory(filepath, reponame)
	if err != nil { 
		if err.Error() == "rpc error: code = Unavailable desc = transport is closing" {
			c.JSON(400, "Query parameters are not right")
		} else {
			c.JSON(500, err.Error())
		}
		return
	}
	c.JSON(200, history.Commits)
	return
}

func GetContentInCommit(c *gin.Context) {
	commit := c.Query("commit")
	filepath := c.Query("file")
	reponame := c.Query("db")
	content, err := grpcclient.GetContentInCommit(commit, filepath, reponame)
	if err != nil {
		if err.Error() == "rpc error: code = Unavailable desc = transport is closing" {
			c.JSON(400, "Query parameters are not right")
		} else {
			c.JSON(500, err.Error())
		}
		return
	}
	c.JSON(200, content.Content)
	return
}

func GetDiffTwoCommitsInFile(c *gin.Context) {
	commitOld := c.Query("commitold")
	commitNew := c.Query("commitnew")
	filepath := c.Query("file")
	reponame := c.Query("db")
	diff, err := grpcclient.GetDiffTwoCommitsInFile(commitOld, commitNew, filepath, reponame)
	if err != nil {
		if err.Error() == "rpc error: code = Unavailable desc = transport is closing" {
			c.JSON(400, "Query parameters are not right")
		} else {
			c.JSON(500, err.Error())
		}
		return
	}
	c.JSON(200, diff.Html)
}

func GetRecordsByQuery(c *gin.Context) {
	methodMsg := "GetRecordsByQuery"
	utils.PrintLogInfo(componentMessage, methodMsg, utils.Query_byquery_ok)
	dbName := c.Param("dbname")
	// body, err := ioutil.ReadAll(c.Request.Body)
    // if err != nil {
    //     utils.PrintLogError(err, componentMessage, methodMsg, "Error reading body")
    //     c.JSON(400, "Query in body cannot be read")
    // }

	var criteriaSet utils.CriteriaSet
	unmarshallingError := c.BindJSON(&criteriaSet)
	if unmarshallingError != nil {
		utils.PrintLogWarn(unmarshallingError, componentMessage, methodMsg, utils.Query_byquery_error)
		c.JSON(400, "Data in body has not expected structure")
		return
	}
	recordSet, err := grpcclient.GetRecordsFromQuery(dbName, criteriaSet)
	if err != nil {
        utils.PrintLogError(err, componentMessage, methodMsg, "Error reading RDB")
        c.JSON(503, "Query sent to server with an error: " + err.Error())
    }
	if len(recordSet.Records) > 0 {
		var recordSetAsStringArray []interface{}
		for _, recordAsSTring := range recordSet.Records {
			var obj interface{}
			err = json.Unmarshal([]byte(recordAsSTring), &obj)
			if err != nil {
				c.JSON(503, "Error converting data")
				return
			}
			recordSetAsStringArray = append(recordSetAsStringArray, obj)
		}
		c.JSON(200, recordSetAsStringArray)
		return
	} else {
		c.JSON(404, "Records not found")
		return
	}
}

/*
Steps:
1-create branch in git useing gitea api
2-create session in admin db
3-return session data

Sending new events to API in session will require the active session id.
It will be checked as active, then using the linked branch in git
*/
func HandleNewSession(c *gin.Context) {
	methodMsg := "HandleNewSession"
	repoName := c.Query("dbname")
	description := c.Query("description")
	user := c.Query("user")
	colID := c.Query("collectionID")
	if !(len(repoName)>0) || !(len(description)>0) || !(len(user)>0) || !(len(colID)>0){
		err := errors.New("Required data is not complete")
		utils.PrintLogError(err, componentMessage, methodMsg, "Error creating session")
		c.JSON(400, "Required data is not complete")
		return
	}
	branchName := utils.GetUUID()
	// Create new branch in Git repo
	s, err := gitapiclient.StartSession(repoName, branchName)
	if err != nil || !(len(s)>0){
		utils.PrintLogError(err, componentMessage, methodMsg, "Error creating session")
		c.JSON(503, "Session has not been created")
		return
	}
	var session utils.Session
	session.Branch = branchName
	session.Description = description
	session.StartTime = 0
	session.EndTime = 0
	session.User = user
	collectionID, err := strconv.Atoi(colID)
	if err != nil {
		c.JSON(400, "Linked cpollection ID is not valid")
		return
	}
	session.Collection.CollectionID = collectionID
	// Record the neew session in the admin DB
	recordedSesssion, recordedSessionErr := admindb.NewSession(session)
	if recordedSessionErr != nil {
		utils.PrintLogError(recordedSessionErr, componentMessage, methodMsg, utils.Session_creation_failure)
		c.JSON(503, recordedSessionErr)
		return
	}
	utils.PrintLogInfo(componentMessage, methodMsg, utils.Session_creation_failure)
	c.JSON(201, recordedSesssion)
	return
}

/*
1-update session in admin db
2-create pull request for the session branch in git useing gitea api
3-return error
*/
func HandleEndSession(c *gin.Context) {
	methodMsg := "HandleEndSession"
	var session utils.Session
	repoName := c.Query("dbname")
	if !(len(repoName)>0){
		err := errors.New("DB name is empty")
		utils.PrintLogError(err, componentMessage, methodMsg, "Error creating session")
		c.JSON(400, "DB name not included")
		return
	}
	unmarshallingError := c.BindJSON(&session)
	if unmarshallingError != nil {
		utils.PrintLogWarn(unmarshallingError, componentMessage, methodMsg, "Error binding request body")
		c.JSON(400, "Data in body has not expected structure")
		return
	}
	commitErr := gitapiclient.CommitSession(repoName, session)
	if commitErr != nil {
		c.JSON(503, "Error ending session in Operational DB - Reason: " + commitErr.Error())
		return
	}
	recordErr := admindb.CloseSession(session.SessionID)
	if recordErr != nil {
		c.JSON(503, "Error ending session in admin DB - Reason: " + recordErr.Error())
		return
	}
	c.JSON(200, fmt.Sprintf("Session ID '%d' committed successfully", session.SessionID))
	return
}

func HandlePostPutRecord(c *gin.Context) {
	methodMsg := "HandlePostPutRecord"
	var event utils.RecordEvent

	utils.PrintLogInfo(componentMessage, methodMsg, utils.Event_received_pre_ok)

	unmarshallingError := c.BindJSON(&event)
	if unmarshallingError != nil {
		utils.PrintLogWarn(unmarshallingError, componentMessage, methodMsg, utils.Event_post_error_client)
		c.JSON(400, "Data in body has not expected structure")
		return
	}
	utils.PrintLogInfo(componentMessage, methodMsg, utils.Event_post_ok)
	fmt.Println("handled event")
	fmt.Println(event)
	if err := handleIncomingRecord(event); err != nil {
		utils.PrintLogError(err, componentMessage, methodMsg, utils.Event_post_error_server)
		c.JSON(400, err.Error())
	}

	utils.PrintLogInfo(componentMessage, methodMsg, utils.Event_post_ok)
	
	c.JSON(202, "Accepted record. This does not mean the record is already written. Please verify later.")
}


func HandlePostPutBatch(c *gin.Context) {
	methodMsg := "HandlePostPutBatch"
	var batch utils.RecordEventBatch

    utils.PrintLogInfo(componentMessage, methodMsg, utils.Event_batch_received_pre_ok)

	unmarshallingError := c.Bind(&batch)
	if unmarshallingError != nil {
		utils.PrintLogWarn(unmarshallingError, componentMessage, methodMsg, "Error unmarshalling batch")
		c.JSON(400, "Data in request body has not expected structure: Record event batch")
		return
	}

	if len(batch.Records) < 1 {
		msg := "The batch of record events is empty"
		utils.PrintLogWarn(errors.New(msg), componentMessage, methodMsg, "")
		c.JSON(400, msg)
		return
	}

	utils.PrintLogInfo(componentMessage, methodMsg, fmt.Sprintf("Processing batch with %d records", len(batch.Records)))
	go handleIncomingRecordBatch(batch)

	utils.PrintLogInfo(componentMessage, methodMsg, utils.Event_batch_received_post_ok)

	c.JSON(202, "Accepted batch. This does not mean the record is already written. Please verify later.")
}

func handleIncomingRecordBatch(batch utils.RecordEventBatch) error {
	methodMsg := "handleIncomingRecordBatch"
	recordEventBatchAsJSON, payloadErr := json.Marshal(batch)
	if payloadErr != nil {
		utils.PrintLogError(payloadErr, componentMessage, methodMsg, "Error marshalling event")
		return payloadErr
	} 
	validationErr := validateRecordEventBatch(batch)
	if validationErr != nil {
		utils.PrintLogError(validationErr, componentMessage, methodMsg, "Batch not valid")
		return validationErr
	} 
	topic := config.Kafka.Githandlebatchtopic
	if !(len(topic) > 0) {
		msg := "Error in configuration - kafka.Githandlebatchtopic (topic name for incoming record batch) property is empty"
		emptyTopicErr := errors.New(msg)
		utils.PrintLogError(emptyTopicErr, componentMessage, methodMsg, msg)
		return nil
	}
	gitTopicErr := kafka.SendMessageToTopic(string(recordEventBatchAsJSON), topic)
	if gitTopicErr != nil {
		utils.PrintLogError(gitTopicErr, componentMessage, methodMsg, fmt.Sprintf("%s - Error sending record event batch to topic '%s'", utils.Event_topic_error, topic))
		return gitTopicErr
	}
	utils.PrintLogInfo(componentMessage, methodMsg, utils.Event_topic_ok)
	return nil
}

func handleIncomingRecord(recordEvent utils.RecordEvent) error {
	methodMsg := "handleIncomingRecord"
	recordEventAsJSON, payloadErr := json.Marshal(recordEvent)
	if payloadErr != nil {
		utils.PrintLogError(payloadErr, componentMessage, methodMsg, "Error marshalling event")
		return payloadErr
	} 
	validationErr := validateRecordEvent(recordEvent)
	if validationErr != nil {
		utils.PrintLogError(validationErr, componentMessage, methodMsg, "Event not valid")
		return validationErr
	} 
	topic := config.Kafka.Githandlerecordtopic
	if !(len(topic) > 0) {
		msg := "Error in configuration - kafka.githandlerecordtopic (topic name for incoming records) property is empty"
		emptyTopicErr := errors.New(msg)
		utils.PrintLogError(emptyTopicErr, componentMessage, methodMsg, msg)
		return nil
	}
	gitTopicErr := kafka.SendMessageToTopic(string(recordEventAsJSON), topic)
	if gitTopicErr != nil {
		utils.PrintLogError(gitTopicErr, componentMessage, methodMsg, fmt.Sprintf("%s - Error sending record event to topic '%s'", utils.Event_topic_error, topic))
		return gitTopicErr
	}
	utils.PrintLogInfo(componentMessage, methodMsg, utils.Event_topic_ok)
	return nil
}

func DeleteRecord(c *gin.Context) {
	methodMsg := "DeleteRecord"
	var event utils.RecordEvent

	unmarshallingError := c.Bind(&event)
	if unmarshallingError != nil {
		utils.PrintLogWarn(unmarshallingError, componentMessage, methodMsg, "Error unmarshalling event")
		c.JSON(400, "Data in body has not expected structure")
		return
	}

	recordAsJSON, payloadErr := json.Marshal(event.RecordContent)
	if payloadErr != nil {
		utils.PrintLogError(payloadErr, componentMessage, methodMsg, fmt.Sprintf( "%s - Error marshalling delete event", utils.Event_delete_error_client))
		c.JSON(503, payloadErr.Error())
	}
	kafkaErr := kafka.SendMessageToTopic(string(recordAsJSON), config.Kafka.Gitdeletetopic)
	if kafkaErr != nil {
		utils.PrintLogError(kafkaErr, componentMessage, methodMsg, fmt.Sprintf( "%s - Error sending delete event to Kafka", utils.Event_topic_delete_error) )
		c.JSON(503, kafkaErr.Error())
	}
	utils.PrintLogInfo(componentMessage, methodMsg, utils.Event_delete_ok)
	c.JSON(202, "Accepted set of records to delete. This does not mean the record is already deleted. Please verify later.")
	return
}

// func DeleteBatch(c *gin.Context) {
// 	methodMsg := "DeleteRecords"
// 	var eventCollection utils.RecordSet

// 	unmarshallingError := c.Bind(&eventCollection)
// 	if unmarshallingError != nil {
// 		utils.PrintLogWarn(unmarshallingError, componentMessage, methodMsg, "Error unmarshalling event")
// 		c.JSON(400, "Data in body has not expected structure")
// 		return
// 	}

// 	if len(eventCollection.Records) < 1 {
// 		msg := "The collection of records is empty"
// 		utils.PrintLogWarn(errors.New(msg), componentMessage, methodMsg, "")
// 		c.JSON(400, msg)
// 		return
// 	}
// 	var mistakes []string
// 	for _, record := range eventCollection.Records {
// 		recordAsJSON, payloadErr := json.Marshal(record)
// 		if payloadErr != nil {
// 			utils.PrintLogError(payloadErr, componentMessage, methodMsg, "Error marshalling delete event")
// 			mistakes = append(mistakes, fmt.Sprintf("Error - Record ID: '%s' - Reason: %s", record.Id, payloadErr.Error))
// 		}
// 		kafkaErr := kafka.SendEventMessageToTopic(string(recordAsJSON), config.Kafka.Gitdeletetopic)
// 		if kafkaErr != nil {
// 			utils.PrintLogError(kafkaErr, componentMessage, methodMsg, "Error sending delete event to Kafka")
// 			mistakes = append(mistakes, fmt.Sprintf("Error - Record ID: '%s' - Reason: %s", record.Id, kafkaErr.Error))
// 		}
// 	}
// 	if len(mistakes) > 0 {
// 		msg := strings.Join(mistakes, " | ")
// 		c.JSON(500, msg)
// 		return
// 	}
// 	c.JSON(200, "Accepted set of records to delete. This does not mean the record is already deleted. Please verify later.")
// 	return
// }

func validateRecordEvent(event utils.RecordEvent) error {
	if len(event.Id) < 1 {
		return errors.New("ID field is empty")
	}
	if len(event.DBName) < 1 {
		return errors.New("DNName field is empty")
	}
	return nil
}
//

func validateRecordEventBatch(batch utils.RecordEventBatch) error {
	if len(batch.Id) < 1 {
		return errors.New("ID field is empty")
	}
	if len(batch.DBName) < 1 {
		return errors.New("DNName field is empty")
	}
	return nil
}