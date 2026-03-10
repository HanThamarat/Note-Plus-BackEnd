package usecase

import (
	"fmt"

	"github.com/HanThamarat/NOTE-PLUS-MESSAGE-QUEUE/internal/domain"
)

type NotifyUsecase struct {
	Listener domain.QueueListener
}

func NewNotifyUsecase (l domain.QueueListener) *NotifyUsecase {
	return &NotifyUsecase{Listener: l}
}

func (u *NotifyUsecase) Start() {
	msgs, err := u.Listener.Consume("order_queue");
	if err != nil {
		fmt.Println("Error starting consumer:", err);
        return;
	}

	fmt.Println(" [*] Waiting for messages. To exit press CTRL+C");

	for msg := range msgs {
		u.process(msg);
	}
}

func (u *NotifyUsecase) process(data []byte) {
    // BUSINESS LOGIC GOES HERE
    // e.g., Send an Email, Update a Database, etc.
    fmt.Printf("Message Received and Processed: %s\n", string(data));
}