package models

import (
	"github.com/streadway/amqp"
)

type DeadQueue struct {
	Original     Queue
	Name         string
	Exchange     string
	ExchangeType string
	RoutingKey   string
	Messages     int
	Durable      bool
	AutoDelete   bool
	Exclusive    bool
	NoWait       bool
	Args         *amqp.Table
}

func NewDeadQueue(original Queue, exchange, exchangeType, name, routingKey string) *DeadQueue {
	queue := DeadQueue{
		Original:     original,
		Name:         name,
		Exchange:     exchange,
		ExchangeType: exchangeType,
		RoutingKey:   routingKey,
		Messages:     0,
		Durable:      true,
		AutoDelete:   false,
		Exclusive:    false,
		NoWait:       false,
		Args:         nil,
	}

	return &queue
}
