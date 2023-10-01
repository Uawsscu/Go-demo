package kafkautil

import (
	"fmt"
	"go-demo/config"

	"os"
)

func SendMessageCreateAnime(message string) error {
	fmt.Println("[SendKafkaMessage] topic:CREATE_KAFKA_MSG start....")
	err := config.SendMessage(os.Getenv("KAFKA_TOPIC_DEMO_CREATE"), message)
	if err != nil {
		return err
	}
	fmt.Println("[SendKafkaMessage] topic:CREATE_KAFKA_MSG success....")
	return nil
}

func SendMessageUpdateAnime(message string) error {
	fmt.Println("SendMessageUpdateAnime start....")
	err := config.SendMessage(os.Getenv("KAFKA_TOPIC_DEMO_UPDATE"), message)
	if err != nil {
		return err
	}
	return nil
}
