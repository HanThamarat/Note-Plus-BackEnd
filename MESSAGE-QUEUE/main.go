package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/HanThamarat/NOTE-PLUS-MESSAGE-QUEUE/internal/domain"
	rabbltmq "github.com/HanThamarat/NOTE-PLUS-MESSAGE-QUEUE/internal/infrastructure/rabbitmq"
	"github.com/HanThamarat/NOTE-PLUS-MESSAGE-QUEUE/internal/usecase"
	"github.com/HanThamarat/NOTE-PLUS-MESSAGE-QUEUE/pkg/database"
	pkg "github.com/HanThamarat/NOTE-PLUS-MESSAGE-QUEUE/pkg/load-env"
)

func main() {
	pkg.LoadEnv();

	db := database.InitDB();

	db.AutoMigrate(
		domain.Project{},
	)

	rabbit, err := rabbltmq.NewRabbitClient();

    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }

    defer rabbit.Close();
	projectWorker := usecase.NewProjectUsecase(rabbit, db);

    go projectWorker.ProjectCreateStartWorker();
	go projectWorker.FindAllProjectByOrg();

	// Create a channel to listen for OS signals (like Ctrl+C or Docker stop)
    stop := make(chan os.Signal, 1);
    signal.Notify(stop, os.Interrupt, syscall.SIGTERM);

    // Wait here until a signal is received
    <-stop

    fmt.Println("Shutting down gracefully...")
}