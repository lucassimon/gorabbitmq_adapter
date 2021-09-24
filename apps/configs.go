package rabbitmq

import (
	"fmt"
	"log"
	"os"
)

type RabbitMQConfig struct {
	user     string
	password string
	host     string
	vhost    string
	dsn      string
}

func NewRabbitMQConfig() *RabbitMQConfig {
	var cfg RabbitMQConfig
	cfg.user = os.Getenv("RABBITMQ_USER")
	cfg.password = os.Getenv("RABBITMQ_PASS")
	cfg.host = os.Getenv("RABBITMQ_HOST")
	cfg.vhost = os.Getenv("RABBITMQ_VHOST")

	cfg.dsn = fmt.Sprintf("amqp://%s:%s@%s/%s", cfg.user, cfg.password, cfg.host, cfg.vhost)

	return &cfg
}

func (cfg *RabbitMQConfig) Dsn() string {
	log.Println("Get the uri", cfg.dsn)
	return cfg.dsn
}
