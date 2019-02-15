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
	url       string
	routeKey  string
	queueName string
	conn      *amqp.Connection
	channel   *amqp.Channel
	err       error
}

func NewRabbitMqClient(url, routeKey, queueName string) (*RabbitMq) {
	client := &RabbitMq{url: url, routeKey: routeKey, queueName: queueName, conn: nil, channel: nil, err: nil}
	err := client.initConnections()

	if err != nil {
		fmt.Println("init rabbit client failed:",err.Error())
		return nil
	}
	return client

}

func (r *RabbitMq) initConnections() (error) {
	r.conn, r.err = amqp.Dial(r.url)

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
	return r.channel.Publish(r.routeKey, r.queueName, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        message})
}

func (r *RabbitMq) Receiver(queueName string) {
	var delivery <-chan amqp.Delivery
	delivery, r.err = r.channel.Consume(r.queueName, "", true, false, false, false, nil)
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
