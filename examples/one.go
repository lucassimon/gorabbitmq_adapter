package examples

import (
	"log"

	rabbitmq "github.com/lucassimon/gorabbitmq_adapter/apps"
	"github.com/lucassimon/gorabbitmq_adapter/apps/interfaces"
	"github.com/lucassimon/gorabbitmq_adapter/apps/models"
	"github.com/streadway/amqp"
)

type DeadQueueService struct {
	Repository interfaces.DeadQueueRepository
}

func NewQueueService(repository interfaces.DeadQueueRepository) *DeadQueueService {
	return &DeadQueueService{
		Repository: repository,
	}
}

func (c *DeadQueueService) CountMessages(channel *amqp.Channel) (*models.DeadQueue, error) {
	log.Println("Services count messages")

	log.Println("Start get messages")
	queue, err := c.Repository.CountAll(channel)

	return queue, err
}

type DeadQueueRepository struct {
	queue   *models.DeadQueue
	Channel *amqp.Channel
}

func NewDeadQueueRepository(queue *models.DeadQueue) *DeadQueueRepository {
	return &DeadQueueRepository{
		queue: queue,
	}
}

func (d *DeadQueueRepository) Count() error {
	log.Println("Counting one queue messages")
	return nil
}

func (d *DeadQueueRepository) Consume() error {
	log.Println("Counting one queue messages")
	return nil
}

func (d *DeadQueueRepository) Publish() error {
	log.Println("Counting one queue messages")
	return nil
}

func (d *DeadQueueRepository) CountAll(channel *amqp.Channel) (*models.DeadQueue, error) {
	log.Println("Counting messages")

	log.Println("exchange: ", d.queue.Exchange, "queue: ", d.queue.Name, d.queue.RoutingKey)

	queue, err := rabbitmq.DeclareQueueInspect(channel, d.queue.Name)

	if err != nil {
		log.Fatalln("Erro when declare queue passive", err)
	}

	d.queue.Messages = queue.Messages

	return d.queue, nil
}

func example_one() {

	log.Println("create rabbitmq configs")
	rabbitMQConfig := rabbitmq.NewRabbitMQConfig()

	log.Println("got Connection, getting Channel")
	conn, err := rabbitmq.RabbitMQConnection(rabbitMQConfig)
	if err != nil {
		log.Fatalf("Error when create a channel: %s", err)
	}
	defer conn.Close()

	channel, err := rabbitmq.RabbitMQCreateChannel(conn)
	if err != nil {
		log.Fatalf("Error when create a channel: %s", err)
	}
	defer channel.Close()
	log.Println("Channel created", channel)

	var args = make(amqp.Table)

	args["x-dead-letter-exchange"] = "DEADQUEUE_EXCHANGE"
	args["x-dead-letter-routing-key"] = "DEADQUEUE_EXAMPLE_RK"

	queue := models.Queue{
		Name:         "EXAMPLE",
		Exchange:     "amqp.direct",
		ExchangeType: models.EXCHANGE_TYPE_DIRECT,
		RoutingKey:   "EXAMPLE_RK",
		Messages:     0,
		Durable:      true,
		AutoDelete:   false,
		Exclusive:    false,
		NoWait:       false,
		Args:         &args,
	}

	deadQueue := models.NewDeadQueue(
		queue,
		"amqp.direct",
		models.EXCHANGE_TYPE_DIRECT,
		"EXAMPLE_DEAD_QUEUE",
		"DEADQUEUE_EXAMPLE_RK",
	)

	log.Println("create a repository")
	repo := NewDeadQueueRepository(deadQueue)

	log.Println("call the service")
	service := NewQueueService(repo)

	log.Println("count messages by queue")
	queueWithMessage, err := service.CountMessages(channel)

	if err != nil {
		log.Fatal("Error when fetch messages", err)
	}

	log.Println(queueWithMessage.Name, "have", queueWithMessage.Messages, "message(s)")

}
