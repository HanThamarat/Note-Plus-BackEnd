package rabbltmq

import (
	"fmt"
	"os"

	"github.com/HanThamarat/NOTE-PLUS-MESSAGE-QUEUE/internal/domain"
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
        return nil, err
    }

    ch, err := conn.Channel()
    if err != nil {
        return nil, err
    }

    return &RabbitClient{Conn: conn, Channel: ch}, nil
}

func (r *RabbitClient) Consume(queueName string) (<-chan domain.Message, error) {
    // 1. Ensure queue exists
    q, err := r.Channel.QueueDeclare(queueName, true, false, false, false, nil)
    if err != nil {
        return nil, err
    }

    msgs, err := r.Channel.Consume(q.Name, "", false, false, false, false, nil)
    if err != nil {
        return nil, err
    }

    out := make(chan domain.Message)
    go func() {
        for d := range msgs {
            out <- domain.Message{
                Body:          d.Body,
                ReplyTo:       d.ReplyTo,
                CorrelationId: d.CorrelationId,
            }

			d.Ack(false);
        }
		close(out);
    }()

    return out, nil
}

func (r *RabbitClient) Publish(queueName string, correlationId string, body []byte) error {
    return r.Channel.Publish("", queueName, false, false, amqp.Publishing{
        ContentType:   "application/json",
        CorrelationId: correlationId,
        Body:          body,
    })
}

func (r *RabbitClient) Close() error {
    r.Channel.Close()
    return r.Conn.Close()
}