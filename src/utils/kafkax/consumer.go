package kafkax

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
)

func Consumer(topic string, partition int) ([]byte, error) {
	// to consume messages
	server := viper.GetString(`kafka.serverAddress`)
	port := viper.GetString(`kafka.serverPort`)

	conn, err := kafka.DialLeader(context.Background(), "tcp", fmt.Sprintf("%s:%s", server, port), topic, partition)
	if err != nil {
		return nil, fmt.Errorf("failed to dial leader: %s", err.Error())
	}

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))

	b := make([]byte, 10e3) // 10KB max per message
	for {
		_, err := conn.Read(b)
		if err != nil {
			return nil, fmt.Errorf("read msg failed: %s", err.Error())
		}
		if len(b) > 0 {
			break
		}
	}

	if err := conn.Close(); err != nil {
		return nil, fmt.Errorf("failed to close connection: %s", err)
	}
	return b, nil
}

// func Consumer(c *kafka.Consumer, topic string) *kafka.Message {

// 	// A signal handler or similar could be used to set this to false to break the loop.
// 	c.SubscribeTopics([]string{"myTopic", "^aRegex.*[Tt]opic"}, nil)

// 	for {
// 		msg, err := c.ReadMessage(time.Second)
// 		if err == nil {
// 			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
// 			return msg
// 		} else if err.(kafka.Error).IsFatal() {
// 			// The client will automatically try to recover from all errors.
// 			// Timeout is not considered an error because it is raised by
// 			// ReadMessage in absence of messages.
// 			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
// 			return nil
// 		}
// 	}
// }
