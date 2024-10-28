package producers

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func TestProducer() {
	con, err := amqp.Dial("amqp://admin:password@localhost:5673/")

	if err != nil {
		log.Fatalln("Error while creating connect to queue")
	}

	defer con.Close()

	ch, err := con.Channel()

	if err != nil {
		log.Fatalln("Error while create rabbitmq channel")
	}

	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	if err != nil {
		log.Fatalf("failed to declare a queue. Error: %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	body := "Hello World!"
	err = ch.PublishWithContext(ctx,
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatalf("failed to publish a message. Error: %s", err)
	}

	log.Printf(" [x] Sent %s\n", body)
}
