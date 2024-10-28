package consumers

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func TestConsumer() {
	con, err := amqp.Dial("amqp://admin:password@localhost:5673/")

	if err != nil {
		log.Fatalln("Error while creating connect to queue")
	}

	defer con.Close()

	ch, err := con.Channel()

	if err != nil {
		log.Fatalln("Error while create rabbitmq channel")
	}

	q, err := ch.QueueDeclare(
		"hello", // имя очереди
		false,   // не динамическая
		false,   // не будет удалена при закрытии
		false,   // не исключительная
		false,   // без ожидания
		nil,     // аргументы
	)

	defer ch.Close()

	messages, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("failed to register a consumer. Error: %s", err)
	}

	var forever chan struct{}
	go func() {
		for message := range messages {
			log.Printf("received a message: %s", message.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
