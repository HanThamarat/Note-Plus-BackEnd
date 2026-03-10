package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	rabbltmq "github.com/HanThamarat/NOTE-PLUS-MESSAGE-QUEUE/internal/infrastructure/rabbitmq"
	"github.com/HanThamarat/NOTE-PLUS-MESSAGE-QUEUE/internal/usecase"
	pkg "github.com/HanThamarat/NOTE-PLUS-MESSAGE-QUEUE/pkg/load-env"
)

func main() {
	pkg.LoadEnv();

	rabbit, err := rabbltmq.NewRabbitClient();

    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }

    defer rabbit.Close();

	notifyWorker := usecase.NewNotifyUsecase(rabbit);
	projectWorker := usecase.NewProjectUsecase(rabbit);

	go notifyWorker.Start();
    go projectWorker.Start();

	// Create a channel to listen for OS signals (like Ctrl+C or Docker stop)
    stop := make(chan os.Signal, 1);
    signal.Notify(stop, os.Interrupt, syscall.SIGTERM);

    // Wait here until a signal is received
    <-stop

    fmt.Println("Shutting down gracefully...")
}