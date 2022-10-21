package kafka

import (
	"time"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	utils "xqledger/apirouter/utils"
	kafka "github.com/segmentio/kafka-go"
	pb "xqledger/apirouter/protobuf"
)

const componentMessage = "Topics Consumer Service"


func getKafkaReader(topic string) *kafka.Reader {
	broker := config.Kafka.Bootstrapserver
	brokers := strings.Split(broker, ",")
	groupID := config.Kafka.Groupid
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MaxWait: 100 * time.Millisecond,
	})
}

func StartListeningForStream(stream pb.RDBQueryService_GetRDBRecordsStreamServer) {
	methodMsg := "StartListeningForStream"
	reader := getKafkaReader(config.Kafka.Gitactionbacktopic)
	defer reader.Close()
	for {
		//m, err := reader.ReadMessage(context.Background())
		m, err := reader.FetchMessage(context.Background()) // explicit commit
		if err != nil {
			utils.PrintLogError(err, componentMessage, methodMsg, fmt.Sprintf("Error reading message - Reason: %s", err.Error()))
		}
		msg := fmt.Sprintf("Message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		utils.PrintLogInfo(componentMessage, methodMsg, msg)
		event, eventErr := convertMessageToProcessable(m)
		if eventErr == nil {
			utils.PrintLogInfo(componentMessage, methodMsg, fmt.Sprintf("Message converted to event successfully - Key '%s'", m.Key))
			recordSet := pb.RecordSet{}
			var records []string
			records[0] = event.RecordContent
			recordSet.Records = records
			sendErr := stream.Send(&recordSet)
			if sendErr != nil {
				utils.PrintLogError(eventErr, componentMessage, methodMsg, fmt.Sprintf("Send output message to stream failed - Reason '%s'", sendErr.Error()))
			}
		}
	}
}


func convertMessageToProcessable(msg kafka.Message) (utils.RecordEvent, error) {
	methodMsg := "convertMessageToProcessable"
	var newRecordEvent utils.RecordEvent
	unmarshalErr := json.Unmarshal(msg.Value, &newRecordEvent)
	if unmarshalErr != nil {
		utils.PrintLogWarn(unmarshalErr, componentMessage, methodMsg, fmt.Sprintf("Error unmarshaling message content to JSON - Key '%s'", msg.Key))
		return newRecordEvent, unmarshalErr
	}
	utils.PrintLogInfo(componentMessage, methodMsg, fmt.Sprintf("ID '%s'", newRecordEvent.Id))
	utils.PrintLogInfo(componentMessage, methodMsg, fmt.Sprintf("Group '%s'", newRecordEvent.Group))
	utils.PrintLogInfo(componentMessage, methodMsg, fmt.Sprintf("OperationType '%s'", newRecordEvent.OperationType))
	return newRecordEvent, nil
}