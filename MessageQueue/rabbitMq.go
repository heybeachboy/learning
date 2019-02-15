package MessageQueue

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

const (
	QueueName = "send"
	Exchange  = "order.send"
	MqUrl     = "amqp://admin:7q4a1z@192.168.0.249:5672/"
)

type RabbitMq struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	err     error
}

func (r *RabbitMq) InitConnections() (error) {
	r.conn, r.err = amqp.Dial(MqUrl)

	if r.err != nil {
		return r.returnError("connection amqp error")
	}

	r.channel, r.err = r.conn.Channel()

	if r.err != nil {
		return r.returnError("open channel error")
	}

	return nil
}

func (r *RabbitMq) Push(message []byte) (error) {
	return r.channel.Publish("", QueueName, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        message})
}

func (r *RabbitMq) Receiver(queueName string) {
	var delivery <-chan amqp.Delivery
	delivery, r.err = r.channel.Consume(queueName, "", true, false, false, false, nil)
	r.HandleDelivery(delivery)
}

func (r *RabbitMq) HandleDelivery(delivery <-chan amqp.Delivery) {
	for msg := range delivery {
		fmt.Println("receive ok:", string(msg.Body))
		msg.Ack(true)
	}

}

func (r *RabbitMq) returnError(message string) (error) {
	return errors.New(fmt.Sprintf("%s:%v", message, r.err))
}
