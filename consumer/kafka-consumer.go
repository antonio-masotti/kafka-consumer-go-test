package consumer

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

// Of course to be retrieved from a config file / env variable / etc.

func ReadFromKafka(kafkaServer string, topic string, offset int64) []kafka.Message {
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

	r.FetchMessage(context.Background())
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
	return messages

}
