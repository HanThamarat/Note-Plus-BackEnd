package usecase

import (
	"fmt"

	"github.com/HanThamarat/NOTE-PLUS-MESSAGE-QUEUE/internal/domain"
)

type ProjectUsecase struct {
	Listener domain.QueueListener
}

func NewProjectUsecase(l domain.QueueListener) *ProjectUsecase {
	return &ProjectUsecase{Listener: l}
}

func (u *ProjectUsecase) Start() {
	msgs, err := u.Listener.Consume("project");
	if err != nil {
		fmt.Println("Error starting consumer:", err);
        return;
	}

	fmt.Println(" [*] Waiting for project messages. To exit press CTRL+C");

	for msg := range msgs {
		u.process(msg);
	}
}

func (u *ProjectUsecase) process(data []byte) {
    // BUSINESS LOGIC GOES HERE
    // e.g., Send an Email, Update a Database, etc.
    fmt.Printf("Project Message Received and Processed: %s\n", string(data));
}