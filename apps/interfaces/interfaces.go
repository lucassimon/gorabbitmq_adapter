package interfaces

import (
	"github.com/lucassimon/gorabbitmq_adapter/apps/models"

	"github.com/streadway/amqp"
)

type QueueRepository interface {
	Count() error
	Consume() error
	Publish() error
}

type DeadQueueRepository interface {
	QueueRepository
	CountAll(channel *amqp.Channel) (*models.DeadQueue, error)
}
