package rabbitmq

import (
	"os"

	error_utils "github.com/MauricioMilano/stock_app/utils/error"

	amqp "github.com/rabbitmq/amqp091-go"
)

var br Broker

func InitilizeBroker() (*amqp.Connection, *amqp.Channel) {

	rmqHost := os.Getenv("RMQ_HOST")
	rmqUserName := os.Getenv("RMQ_USERNAME")
	rmqPassword := os.Getenv("RMQ_PASSWORD")
	rmqPort := os.Getenv("RMQ_PORT")
	dsn := "amqp://" + rmqUserName + ":" + rmqPassword + "@" + rmqHost + ":" + rmqPort + "/"

	conn, err := amqp.Dial(dsn)
	error_utils.ErrorCheck(err)

	ch, err := conn.Channel()
	error_utils.ErrorCheck(err)

	br.SetUp(ch)
	return conn, ch
}

func GetRabbitMQBroker() *Broker {
	return &br
}
