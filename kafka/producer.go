package kafka

import (
	"fmt"
	"context"
	"os"
	"strings"
	"github.com/google/uuid"
	configuration "xqledger/apirouter/configuration"
	utils "xqledger/apirouter/utils"
	kafka "github.com/segmentio/kafka-go"
)

const componentProducerMessage = "Kafka Producer Service"

var config = configuration.GlobalConfiguration

func getKafkaWriter(topic string) *kafka.Writer {
	var listOfBrokersInConfig = os.Getenv("KAFKA_BOOTSTRAPSERVER")
	if !(len(listOfBrokersInConfig) > 0) {
		listOfBrokersInConfig = config.Kafka.Bootstrapserver
	}
	brokers := strings.Split(listOfBrokersInConfig, ",")
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  brokers,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
}


func SendMessageToTopic(msg string, topic string) error {
	methodMsg := "SendMessageToTopic"
	kafkaWriter := getKafkaWriter(topic)
	defer kafkaWriter.Close()
	topicContent := kafka.Message{
		Key:   []byte(uuid.New().String()),
		Value: []byte(msg),
	}
	err := kafkaWriter.WriteMessages(context.Background(), topicContent)
	if err != nil {
		utils.PrintLogError(err, componentProducerMessage, methodMsg, fmt.Sprintf("Error writing message to topic '%s'", topic))
		return err
	}
	utils.PrintLogInfo(componentProducerMessage, methodMsg, fmt.Sprintf("Message sent to topic '%s' successfully", topic)) 

	return nil
}
