/**
* Go File created on 13.01.23
* by Antonio Masotti (antonio)
* MIT License
 */

package producer

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

func CreateWriter(kafkaServer string, topic string) *kafka.Writer {
	w := &kafka.Writer{
		Addr:     kafka.TCP(kafkaServer),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	return w
}

func WriteToKafka(writer *kafka.Writer, messageKey string, messageBody string) {
	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(messageKey),
			Value: []byte(messageBody),
		},
	)

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

}
