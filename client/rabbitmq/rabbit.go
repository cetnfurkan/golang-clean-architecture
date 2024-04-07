package rabbitmq

import (
	"fmt"
	"golang-clean-architecture/config"
	"golang-clean-architecture/repository"
	"golang-clean-architecture/service"
	"golang-clean-architecture/target/ent"

	"github.com/cetnfurkan/core/cache"
	"github.com/cetnfurkan/core/mq"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	config  *config.Config
	db      *ent.Client
	cache   cache.Cache
	service service.UserService
	rabbit  mq.MQ
}

func NewRabbitMQ(config *config.Config, db *ent.Client, cache cache.Cache) *RabbitMQ {

	userRepository := repository.NewUserPostgresRepository(db)

	service := service.NewUserServiceImpl(&config.Echo, userRepository, cache)

	rabbitmq := &RabbitMQ{config: config, db: db, cache: cache, service: service}

	rabbitmq.rabbit = mq.NewRabbitMQ(
		&config.RabbitMQ,
		mq.WithRabbitConsumerMessageHandler(rabbitmq.consumerHandler),
	)

	return rabbitmq
}

func (rabbitmq RabbitMQ) Consume() {

	go func() {
		err := rabbitmq.rabbit.Consumer().Consume("consume_path")
		if err != nil {
			fmt.Println("Error consuming message: ", err)
		}
	}()
}

// func (rabbitmq RabbitMQ) Produce() {

// 	err := rabbitmq.rabbit.(*mq.RabbitMQ).ProducerWith(
// 		producer.WithRabbitMQArgs(
// 			amqp.Table{
// 				"x-dead-letter-exchange":    "",
// 				"x-dead-letter-routing-key": "yoiur_queue_name",
// 			},
// 		),
// 		producer.WithRabbitMQExpiration("10000"),
// 	).Produce("yoiur_queue_name", []byte("test"))
// 	if err != nil {
// 		fmt.Println("Error producing message: ", err)
// 	}
// }

func (rabbitmq RabbitMQ) consumerHandler(msg *amqp.Delivery) {
	fmt.Println("Message consumed: ", string(msg.Body))
}
