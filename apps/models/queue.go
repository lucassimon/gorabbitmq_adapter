package models

import "github.com/streadway/amqp"

type Queue struct {
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

func NewQueue(exchange, exchangeType, name, routingKey string) *Queue {
	queue := Queue{
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
