package rabbitmq

import (
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitClient struct {
    Conn    *amqp.Connection
    Channel *amqp.Channel
}

func NewRabbitClient() (*RabbitClient, error) {
	
    host 		:= os.Getenv("RABBIT_HOST");
	port 		:= os.Getenv("RABBIT_PORT");
	username 	:= os.Getenv("RABBIT_USERNAME");
	password 	:= os.Getenv("RABBIT_PASSWORD");

	dsn := fmt.Sprintf("amqp://%s:%s@%s:%s/", username, password, host, port);
    
    conn, err := amqp.Dial(dsn)
    if err != nil {
        return nil, err;
    }

    ch, err := conn.Channel()
    if err != nil {
        return nil, err;
    }

    return &RabbitClient{Conn: conn, Channel: ch}, nil;
}

func (r *RabbitClient) Publish(queueName string, body []byte) error {
    _, err := r.Channel.QueueDeclare(queueName, true, false, false, false, nil);
    if err != nil {
        return err;
    }

    return r.Channel.Publish("", queueName, false, false, amqp.Publishing{
        ContentType: "text/plain",
        Body:        body,
    });
}

func (r *RabbitClient) Close() error {
    r.Channel.Close();
    return r.Conn.Close();
}