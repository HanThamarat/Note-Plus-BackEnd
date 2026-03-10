package domain

import "encoding/json"

type Message struct {
    Body          []byte
    ReplyTo       string
    CorrelationId string
}

type MessageBodyDTO struct {
    Status        bool
    Body          json.RawMessage
}

type QueueListener interface {
    Consume(qName string) (<-chan Message, error);
    Publish(queueName string, correlationId string, body []byte) error
}