package kafkax

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func (k *KafkaClient) Producer(topic string, message []byte) {

	// Produce messages to topic (asynchronously)
	k.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
	}, nil)

	// Wait for message deliveries before shutting down
	k.P.Flush(15 * 1000)
}

func (k *KafkaClient) WatcherProducer() {
	//Watch Producer Delivery report handler for produced messages
	go func() {
		for e := range k.P.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()
}
