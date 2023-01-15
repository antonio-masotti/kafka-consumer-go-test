package consumer

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

// Of course to be retrieved from a config file / env variable / etc.

// ReadFromKafka reads from a kafka topic
// func ReadFromKafka(kafkaServer string, topic string, offset int64) []kafka.Message {
// 	r := kafka.NewReader(kafka.ReaderConfig{
// 		Brokers:   []string{kafkaServer},
// 		Topic:     topic,
// 		Partition: 0,
// 		MinBytes:  10e3, // 10KB
// 		MaxBytes:  10e6, // 10MB
// 	})
func ReadFromKafka(kafkaServer string, topic string, offset int64) ([]kafka.Message, error) {
	readerConfig := kafka.ReaderConfig{
		Brokers:     []string{kafkaServer},
		Topic:       topic,
		MinBytes:    0,    // 10KB
		MaxBytes:    10e6, // 10MB
		MaxAttempts: 3,
	}
	r := kafka.NewReader(readerConfig)
	if offset > 0 {
		_ = r.SetOffset(offset)
	}

	_, err := r.FetchMessage(context.Background())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	n := r.Stats().Messages
	messages := make([]kafka.Message, n)

	for i := 0; i < int(n-1); i++ {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			panic(err)
		}
		fmt.Printf("Received message nr. %d\n", i)
		messages[i] = m
	}

	r.Close()
	return messages, nil

}
