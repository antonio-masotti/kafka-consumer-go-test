package main

import (
	"fmt"
	"github.com/Bikeleasing-Service/kafka-consumer/api"
	"github.com/Bikeleasing-Service/kafka-consumer/consumer"
	"github.com/Bikeleasing-Service/kafka-consumer/producer"
)

var kafkaServer = "localhost:9092"
var readTopic = "test-topic"
var writeTopic = "test-producer"

func main() {

	// READ
	messages := consumer.ReadFromKafka(kafkaServer, readTopic, 0)

	// Create Writer
	w := producer.CreateWriter(kafkaServer, writeTopic)
	defer w.Close()

	// WRITE to another topic
	for i, m := range messages {
		if m.Value != nil {
			fmt.Printf("Message at offset %d: %s = %s with Headers %s\n", m.Offset, string(m.Key), string(m.Value), m.Headers)
			key := fmt.Sprintf("re-rewritten-message-number-%d", i)
			producer.WriteToKafka(w, key, string(m.Value))
			api.SendToApi(string(m.Value))
		}
	}

}
