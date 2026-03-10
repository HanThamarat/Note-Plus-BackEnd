package domain

import "encoding/json"

type MessageBodyDTO struct {
    Status        bool
    Body          json.RawMessage
}

type QueueProvider interface {
    Publish(queueName string, body []byte) error
	Call(queueName string, body []byte) ([]byte, error)
}