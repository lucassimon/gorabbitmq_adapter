package rabbitmq

import (
	"github.com/streadway/amqp"
)

func RabbitMQConnection(setting *RabbitMQConfig) (*amqp.Connection, error) {
	amqp_uri := setting.Dsn()
	conn, err := amqp.DialConfig(amqp_uri, amqp.Config{Properties: amqp.Table{"connection_name": "test backoffice"}})

	if err != nil {
		return nil, err
	}

	return conn, nil
}

func RabbitMQCreateChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return channel, nil
}
