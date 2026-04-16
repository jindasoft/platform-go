package xmq

import (
	"context"
	"fmt"

	"github.com/jindasoft/jinda-platform/xlogger"
	amqp "github.com/rabbitmq/amqp091-go"
)

const closeChannelErrMsg = "Failed to close channel: %v"

type RabbitMQConfig struct {
	Host        string
	Port        int
	User        string
	Password    string `json:"-"`
	VirtualHost string
}

type RabbitMQService interface {
	Publish(queueName string, body []byte) error
	Consume(queueName string) (<-chan amqp.Delivery, error)
	PublishTopic(exchange, routingKey string, body []byte) error
	ConsumeTopic(exchange, queueName, bindingKey string) (<-chan amqp.Delivery, error)
	Close() error
}

type rabbitMQServiceImpl struct {
	Conn *amqp.Connection
}

func NewRabbitMQ(ctx context.Context, cfg *RabbitMQConfig) (RabbitMQService, error) {
	conn, err := amqp.DialConfig(
		fmt.Sprintf("amqp://%s:%s@%s:%d/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.VirtualHost),
		amqp.Config{},
	)
	if err != nil {
		xlogger.SysErrorf("Failed to connect to RabbitMQ: %v", err)
		return nil, err
	}

	xlogger.SysInfof("RabbitMQ initialized successfully")
	return &rabbitMQServiceImpl{Conn: conn}, nil
}

func (r *rabbitMQServiceImpl) PublishTopic(exchange, routingKey string, body []byte) error {
	ch, err := r.Conn.Channel()
	if err != nil {
		xlogger.SysErrorf("Failed to open channel: %v", err)
		return err
	}
	defer func() {
		if err := ch.Close(); err != nil {
			xlogger.SysErrorf(closeChannelErrMsg, err)
		}
	}()

	err = ch.ExchangeDeclare(
		exchange,
		"topic",
		true,  // durable
		false, // autoDelete
		false, // internal
		false, // noWait
		nil,   // args
	)
	if err != nil {
		xlogger.SysErrorf("Failed to declare exchange: %v", err)
		return err
	}

	return ch.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
}

// Consume messages from a topic exchange
func (r *rabbitMQServiceImpl) ConsumeTopic(exchange, queueName, bindingKey string) (<-chan amqp.Delivery, error) {
	ch, err := r.Conn.Channel()
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		exchange,
		"topic",
		true,  // durable
		false, // autoDelete
		false, // internal
		false, // noWait
		nil,   // args
	)
	if err != nil {
		if closeErr := ch.Close(); closeErr != nil {
			xlogger.SysErrorf(closeChannelErrMsg, closeErr)
		}
		return nil, err
	}

	_, err = ch.QueueDeclare(
		queueName,
		true,  // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,   // args
	)
	if err != nil {
		if closeErr := ch.Close(); closeErr != nil {
			xlogger.SysErrorf(closeChannelErrMsg, closeErr)
		}
		return nil, err
	}

	err = ch.QueueBind(
		queueName,
		bindingKey,
		exchange,
		false,
		nil,
	)
	if err != nil {
		if closeErr := ch.Close(); closeErr != nil {
			xlogger.SysErrorf("Failed to close channel: %v", closeErr)
		}
		return nil, err
	}

	msgs, err := ch.Consume(
		queueName,
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		if closeErr := ch.Close(); closeErr != nil {
			xlogger.SysErrorf(closeChannelErrMsg, closeErr)
		}
		return nil, err
	}
	return msgs, nil
}

func (r *rabbitMQServiceImpl) Publish(queueName string, body []byte) error {
	ch, err := r.Conn.Channel()
	if err != nil {
		return err
	}
	defer func() {
		if err := ch.Close(); err != nil {
			xlogger.SysErrorf(closeChannelErrMsg, err)
		}
	}()

	_, err = ch.QueueDeclare(
		queueName,
		true,  // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,   // args
	)
	if err != nil {
		return err
	}

	return ch.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
}

func (r *rabbitMQServiceImpl) Consume(queueName string) (<-chan amqp.Delivery, error) {
	ch, err := r.Conn.Channel()
	if err != nil {
		return nil, err
	}
	// ไม่ปิด channel ที่นี่ ให้ caller ปิดเองหลังใช้งานเสร็จ

	_, err = ch.QueueDeclare(
		queueName,
		true,  // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,   // args
	)
	if err != nil {
		if closeErr := ch.Close(); closeErr != nil {
			xlogger.SysErrorf(closeChannelErrMsg, closeErr)
		}
		return nil, err
	}

	msgs, err := ch.Consume(
		queueName,
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		if closeErr := ch.Close(); closeErr != nil {
			xlogger.SysErrorf(closeChannelErrMsg, closeErr)
		}
		return nil, err
	}
	return msgs, nil
}

func (r *rabbitMQServiceImpl) Close() error {
	if r.Conn != nil {
		return r.Conn.Close()
	}
	return nil
}
