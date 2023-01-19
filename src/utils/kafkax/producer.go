package kafkax

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
)

func Producer(msg []byte, topic string, partition int) error {

	// to produce messages
	server := viper.GetString(`kafka.serverAddress`)
	port := viper.GetString(`kafka.serverPort`)

	conn, err := kafka.DialLeader(context.Background(), "tcp", fmt.Sprintf("%s:%s", server, port), topic, partition)
	if err != nil {
		return fmt.Errorf("failed to dial leader: %s", err.Error())
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: msg},
	)
	if err != nil {
		return fmt.Errorf("failed to write messages: %s", err.Error())
	}

	if err := conn.Close(); err != nil {
		return fmt.Errorf("failed to close writer: %s", err.Error())
	}

	return nil
}

// func Producer(p *kafka.Producer, topic string, message []byte) {

// 	// Produce messages to topic (asynchronously)
// 	p.Produce(&kafka.Message{
// 		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
// 		Value:          message,
// 	}, nil)

// 	// Wait for message deliveries before shutting down
// 	p.Flush(15 * 1000)
// }

// func WatcherProducer(p *kafka.Producer) {
// 	//Watch Producer Delivery report handler for produced messages
// 	go func() {
// 		for e := range p.Events() {
// 			switch ev := e.(type) {
// 			case *kafka.Message:
// 				if ev.TopicPartition.Error != nil {
// 					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
// 				} else {
// 					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
// 				}
// 			}
// 		}
// 	}()
// }
