package rabbitmq

import (
	"base/gmq/base"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	base.MQ                  // extends Driver, NodeName
	Conn    *amqp.Connection // rabbitmq connection
}

func (r *RabbitMQ) Connect(node string) (err error) {
	conn, err := amqp.Dial(node)
	if err != nil {
		return err
	}
	r.Conn = conn
	r.NodeName = node
	return nil
}

func (r *RabbitMQ) Close() (err error) {
	return r.Conn.Close()
}

func (r *RabbitMQ) Shard(k string) (err error) {
	// if not connect, do sharding
	if r.Conn == nil {
		return r.Connect(r.Driver.GetShardNode(k))
	}
	return nil
}

func (r *RabbitMQ) Publish(qName string, qBody string) (err error) {
	// shard by queue name
	r.Shard(qName)
	// open channel
	ch, err := r.Conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	// open queue
	q, err := ch.QueueDeclare(
		qName, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}
	// publish
	return ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(qBody),
		})
}

func (r *RabbitMQ) Consume(qName string, callback func(done chan bool, body []byte)) (err error) {
	// shard by queue name
	r.Shard(qName)
	// open channel
	ch, err := r.Conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	// open queue
	q, err := ch.QueueDeclare(
		qName, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}
	done := make(chan bool)
	go func() {
		for d := range msgs {
			callback(done, d.Body)
		}
	}()
	<-done // wait for done
	return nil
}
