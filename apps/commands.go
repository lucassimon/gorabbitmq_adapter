package rabbitmq

import (
	"log"

	"github.com/lucassimon/gorabbitmq_adapter/apps/models"

	"github.com/streadway/amqp"
)

func DeclareExchange(ch *amqp.Channel, exchange models.Exchange) error {
	return ch.ExchangeDeclare(
		exchange.Name,        // name
		exchange.Type,        // type
		exchange.Durable,     // durable
		exchange.AutoDeleted, // auto-deleted
		exchange.Internal,    // internal
		exchange.NoWait,      // no-wait
		*exchange.Args,       // arguments
	)
}

func DeclareQueue(ch *amqp.Channel, q models.Queue) (amqp.Queue, error) {
	return ch.QueueDeclare(
		q.Name,       // name
		q.Durable,    // durable
		q.AutoDelete, // delete when unused
		q.Exclusive,  // exclusive
		q.NoWait,     // no-wait
		*q.Args,      // arguments
	)
}

func DeclareQueuePassive(ch *amqp.Channel, q *models.Queue) (amqp.Queue, error) {
	log.Println("Start declare queue passive", ch)

	queue, err := ch.QueueDeclarePassive(
		q.Name,       // name
		q.Durable,    // durable
		q.AutoDelete, // delete when unused
		q.Exclusive,  // exclusive
		q.NoWait,     // no-wait
		*q.Args,      // arguments
	)

	return queue, err
}

func DeclareQueueInspect(ch *amqp.Channel, name string) (amqp.Queue, error) {
	return ch.QueueInspect(name)
}

func DeclareBind(ch *amqp.Channel, queueName, exchange, key string) error {
	return ch.QueueBind(
		queueName, // name of the queue
		key,       // bindingKey
		exchange,  // sourceExchange
		false,     // noWait
		nil,       // arguments
	)
}

func DeclareQos(ch *amqp.Channel) error {
	return ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
}
