package kafkax

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/viper"
)

type KafkaClient struct {
	P *kafka.Producer
	C *kafka.Consumer
}

func InitKafka() (*kafka.Producer, *kafka.Consumer) {

	server := viper.GetString(`kafka.serverAddress`)

	//Producer
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": server})
	if err != nil {
		panic(err)
	}

	//Consumer
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": server,
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}

	return p, c
}
