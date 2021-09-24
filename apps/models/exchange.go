package models

import "github.com/streadway/amqp"

const (
	EXCHANGE_TYPE_DIRECT = "direct"
	EXCHANGE_TYPE_FANOUT = "fanout"
	EXCHANGE_TYPE_TOPIC  = "topic"
)

type Exchange struct {
	Name        string
	Type        string
	Durable     bool
	AutoDeleted bool
	Internal    bool
	NoWait      bool
	Args        *amqp.Table
}
