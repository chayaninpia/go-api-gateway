package kafkax

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
)

func Consumer(ctx context.Context, topic string, partition int) ([]byte, error) {
	// to consume messages
	server := viper.GetString(`kafka.serverAddress`)
	port := viper.GetString(`kafka.serverPort`)

	// conn, err := kafka.DialLeader(context.Background(), "tcp", fmt.Sprintf("%s:%s", server, port), topic, partition)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to dial leader: %s", err.Error())
	// }
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{fmt.Sprintf("%s:%s", server, port)},
		Topic:     topic,
		Partition: partition,
		MaxBytes:  10e5,
	})

	msg := kafka.Message{}
	for {

		m, err := reader.FetchMessage(ctx)
		if err != nil {
			return nil, fmt.Errorf("read msg failed: %s", err.Error())
		}

		if len(m.Value) > 0 {
			reader.CommitMessages(ctx, m)
			msg = m
			break
		}
	}

	if err := reader.Close(); err != nil {
		return nil, fmt.Errorf("failed to close connection: %s", err)
	}

	// var msg string
	// log.Println(string(b))
	// if err := json.Unmarshal(b, &msg); err != nil {
	// 	return "", fmt.Errorf("unmarshal msg failed: %s", err.Error())
	// }
	return msg.Value, nil
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
