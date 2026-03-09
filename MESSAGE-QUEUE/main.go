package main

import (
	"fmt"
	"log"

	errorinside "github.com/HanThamarat/NOTE-PLUS-MESSAGE-QUEUE/pkg/err"
	pkg "github.com/HanThamarat/NOTE-PLUS-MESSAGE-QUEUE/pkg/load-env"
	"github.com/HanThamarat/NOTE-PLUS-MESSAGE-QUEUE/pkg/rabbltmq"
)

func main() {
	pkg.LoadEnv();

	fmt.Println("hello world");

	ch := rabbltmq.RabbitConnection();

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	);

	errorinside.FailOnError(err, "Failed to declare a queue");

	msgs, err := ch.Consume(
	q.Name, // queue
	"",     // consumer
	true,   // auto-ack
	false,  // exclusive
	false,  // no-local
	false,  // no-wait
	nil,    // args
	)
	
	errorinside.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
	}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}