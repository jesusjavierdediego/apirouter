package grpcclient

import (
	"fmt"
	"context"
	"log"
	"time"
	"strconv"
    "google.golang.org/grpc"
	pb "xqledger/apirouter/protobuf"
	utils "xqledger/apirouter/utils"
)

var rdbreader_address = config.Grpcclient.Rdbreaderhost + ":" + strconv.Itoa(config.Grpcclient.Rdbreaderport)
var rdbreader_conn *grpc.ClientConn
var rdbreader_connErr error

func getRDBReaderServerConn() (*grpc.ClientConn, error){
	rdbreader_conn, rdbreader_connErr = grpc.Dial(rdbreader_address, grpc.WithInsecure())
	if rdbreader_connErr != nil {
		log.Fatalf("did not connect: %v", rdbreader_connErr)
		return nil, rdbreader_connErr
	}
	return rdbreader_conn, nil
}

// message Criteria {
//     string boolean_operator = 1; // 'OR' | AND
//     string field = 2;
//     string is = 3; // 'equal' | 'like' 
//     string value = 4;
// }

// message RDBQuery {
//     string database_name = 1;
//     string collection_name = 2;
//     repeated Criteria query = 3;
// }

func GetRecordsFromQuery(dbName string, criteriaSet utils.CriteriaSet) (*pb.RecordSet, error){
	var methodMessage = "GetRecordsFromQuery"
	var emptyResult pb.RecordSet
	rdbreader_conn, rdbreader_connErr = getRDBReaderServerConn()
	if rdbreader_connErr != nil {
		utils.PrintLogError(rdbreader_connErr, componentMessage, methodMessage, "Error in connection")
		return &emptyResult, rdbreader_connErr
	}
	defer rdbreader_conn.Close()
	c := pb.NewRDBQueryServiceClient(rdbreader_conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var rdbQuery []*pb.Criteria
	if len(criteriaSet.Set) > 0 {
		for _, c := range criteriaSet.Set {
			var pbcriteria *pb.Criteria
			pbcriteria.BooleanOperator = c.Booleanoperator
			pbcriteria.Field = c.Parameter
			pbcriteria.Value = c.Value
			pbcriteria.Is = c.Operator
			rdbQuery = append(rdbQuery, pbcriteria)
		}
	}
	
	var q pb.RDBQuery
	q.DatabaseName = dbName
	q.CollectionName = "main" // TODO hardcoded for now, to be decided in API
	q.Query = rdbQuery

	recordSet, err := c.GetRDBRecords(ctx, &q)
	if err != nil {
		utils.PrintLogError(err, componentMessage, methodMessage, "Error in grpc server")
		return &emptyResult, err
	}

	utils.PrintLogInfo(componentMessage, methodMessage, fmt.Sprintf("Number of successfully retrieved records: %d", len(recordSet.Records)))
	return recordSet, nil
}


func GetCountFromColl(dbName, collection string) (int64, error){
	var methodMessage = "GetCountFromColl"
	if !(len(collection)>0){
		collection = "main"
	}
	rdbreader_conn, rdbreader_connErr = getRDBReaderServerConn()
	if rdbreader_connErr != nil {
		utils.PrintLogError(rdbreader_connErr, componentMessage, methodMessage, "Error in connection")
		return 0, rdbreader_connErr
	}
	defer rdbreader_conn.Close()
	c := pb.NewRDBQueryServiceClient(rdbreader_conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var rdbQuery []*pb.Criteria
	
	var q pb.RDBQuery
	q.DatabaseName = dbName
	q.CollectionName = collection
	q.Query = rdbQuery

	collCount, err := c.GetNumberRecordsFromColl(ctx, &q)
	if err != nil {
		utils.PrintLogError(err, componentMessage, methodMessage, "Error in grpc server")
		return 0, err
	}

	utils.PrintLogInfo(componentMessage, methodMessage, fmt.Sprintf("Number of records in collection 'main'-database '%s': %d", dbName, collCount.Count))
	return collCount.Count, nil
}