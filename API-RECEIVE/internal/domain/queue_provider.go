package domain

type QueueProvider interface {
    Publish(queueName string, body []byte) error
}