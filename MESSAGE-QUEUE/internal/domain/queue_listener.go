package domain

type QueueListener interface {
    Consume(qName string) (<-chan []byte, error);
}