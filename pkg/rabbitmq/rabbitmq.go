package rabbitmq

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type RabbitMQConfig struct {
	Host         string
	Port         string
	User         string
	Password     string
	ExchangeName string
	Kind         string
}

// Initialize new channel for rabbitmq
func NewRabbitMQConn(cfg *RabbitMQConfig) (*amqp.Connection, error, func()) {

	connAddr := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
	)

	conn, err := amqp.Dial(connAddr)
	if err != nil {
		log.Error(err, "Failed to connect to RabbitMQ")
		return nil, err, nil
	}

	return conn, nil, func() {
		_ = conn.Close()
	}
}
