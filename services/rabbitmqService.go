package services

import (
	"fmt"

	"github.com/streadway/amqp"
	_ "github.com/streadway/amqp"
)

type RabbitmqService struct {
	Conn *amqp.Connection
	Ch   *amqp.Channel
}

func NewRabbitMQService() *RabbitmqService {
	new_conn := CreateConn()
	new_ch := CreateChannel(*new_conn)
	CreateQueue(*new_ch)
	return &RabbitmqService{Conn: new_conn, Ch: new_ch}
}

func CreateConn() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return conn
}

func CreateChannel(conn amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return ch
}

func CreateQueue(ch amqp.Channel) *amqp.Queue {
	q, err := ch.QueueDeclare(
		"MainQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return &q
}

func (q *RabbitmqService) Publish(message string) {
	err := q.Ch.Publish(
		"",
		"MainQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
