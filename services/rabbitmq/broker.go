package rabbitmq

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/MauricioMilano/stock_app/services/websocket"
	"github.com/MauricioMilano/stock_app/utils"
	error_utils "github.com/MauricioMilano/stock_app/utils/error"

	amqp "github.com/rabbitmq/amqp091-go"
)

type StockRequest struct {
	RoomId uint   `json:"RoomId"`
	Code   string `json:"Code"`
}

type StockResponse struct {
	RoomId  uint   `json:"RoomId"`
	Message string `json:"Message"`
}

type Broker struct {
	ReceiverQueue  amqp.Queue
	PublisherQueue amqp.Queue
	Channel        *amqp.Channel
}

// Setup creates(or connects if not existing) the reciever and publisher queues
func (b *Broker) SetUp(ch *amqp.Channel) {
	receiverQueue := os.Getenv("RECEIVER_QUEUE")
	publisherQueue := os.Getenv("PUBLISHER_QUEUE")

	q1, err := ch.QueueDeclare(
		receiverQueue, // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	error_utils.ErrorCheck(err)

	q2, err := ch.QueueDeclare(
		publisherQueue, // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	error_utils.ErrorCheck(err)

	b.ReceiverQueue = q1
	b.PublisherQueue = q2
	b.Channel = ch
}

func (b *Broker) PublishMessage(requestBody chan []byte) {
	for body := range requestBody {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		err := b.Channel.PublishWithContext(ctx,
			"",                    // exchange
			b.PublisherQueue.Name, // routing key
			false,                 // mandatory
			false,                 // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        body,
			})
		cancel()
		if err != nil {
			log.Printf("PublishMessage Error occured %s\n", err)
			continue
		}
		log.Printf("Sent %s\n", body)
	}
}

func (b *Broker) ReadMessages(roomMap *websocket.RoomMap) {
	msgs, err := b.Channel.Consume(
		b.ReceiverQueue.Name, // queue
		"",                   // consumer
		true,                 // auto-ack
		false,                // exclusive
		false,                // no-local
		false,                // no-wait
		nil,                  // args
	)
	if err != nil {
		log.Printf("ReadMessages Error occured %s\n", err)
		return
	}

	rsvdMsgs := make(chan StockResponse)
	go messageTransformer(msgs, rsvdMsgs)
	go processResponse(rsvdMsgs, b, roomMap)
	select {}
}

func messageTransformer(entries <-chan amqp.Delivery, receivedMessages chan StockResponse) {
	var sr StockResponse
	for d := range entries {
		err := utils.ParseByteArray(d.Body, &sr)
		if err != nil {
			log.Printf("Received bad response : %s ", string(d.Body))
			continue
		}
		log.Println("Received a response")
		receivedMessages <- sr
	}
}

func processResponse(s <-chan StockResponse, b *Broker, room *websocket.RoomMap) {
	for r := range s {
		log.Println("processing stock response for ", r.Message)

		sr := StockResponse{
			RoomId:  r.RoomId,
			Message: r.Message,
		}
		pool := room.GetPool(strconv.FormatUint(uint64(sr.RoomId), 10))
		message := websocket.Message{Type: 3, Body: websocket.Body{ChatRoomId: int32(sr.RoomId), ChatUser: "stock-api-bot", ChatMessage: sr.Message}}
		pool.Broadcast <- message
		log.Println("processed", sr.Message)
	}
}
