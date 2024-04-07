package kafka

import (
	"fmt"
	"log"
	"time"

	"github.com/cetnfurkan/core/config"
	"github.com/cetnfurkan/core/mq"
	"github.com/segmentio/kafka-go"
)

var (
	dialerHandler = mq.WithKafkaDialerHandler(func() (*kafka.Dialer, error) {
		return &kafka.Dialer{
			Timeout:   10 * time.Second,
			DualStack: true,
		}, nil
	})

	msgHandler = mq.WithKafkaConsumerMessageHandler(func(message *kafka.Message) {
		fmt.Println("consumed message", string(message.Value))
	})
)

func Consume(cfg *config.MQ) {
	kafkamq := mq.NewKafka(cfg, dialerHandler, msgHandler)
	err := kafkamq.Consumer().Consume("foo")
	if err != nil {
		log.Fatal(err)
	}
}

func Produce(cfg *config.MQ) {
	kafkamq := mq.NewKafka(cfg, dialerHandler, msgHandler)
	err := kafkamq.Producer().Produce("foo", []byte("bar"))
	if err != nil {
		log.Fatal(err)
	}
}
