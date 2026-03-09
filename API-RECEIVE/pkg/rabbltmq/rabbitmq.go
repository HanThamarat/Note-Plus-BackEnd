package rabbltmq

import (
	"fmt"
	"os"

	errorinside "github.com/HanThamarat/Note-Plus-BackEnd/pkg/err"
	amqp "github.com/rabbitmq/amqp091-go"
)



func RabbitConnection() *amqp.Channel {

	host 		:= os.Getenv("HOST");
	port 		:= os.Getenv("PORT");
	username 	:= os.Getenv("USERNAME");
	password 	:= os.Getenv("PASSWORD");

	credential := fmt.Sprintf("amqp://%s:%s@%s:%s/", username, password, host, port);

	conn, err := amqp.Dial(credential);
  	errorinside.FailOnError(err, "Failed to connect to RabbitMQ");

	ch, err := conn.Channel();
	errorinside.FailOnError(err, "Failed to open a channel");

	return ch;
}