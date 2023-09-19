package kafkautil

import (
	"fmt"
	"main/config"
)

func SendMessageCreateAnime(message string) error {
	fmt.Println("[SendKafkaMessage] topic:CREATE_KAFKA_MSG start....")
	err := config.SendMessage("CREATE_KAFKA_MSG", message)
	if err != nil {
		return err
	}
	fmt.Println("[SendKafkaMessage] topic:CREATE_KAFKA_MSG success....")
	return nil
}

func SendMessageUpdateAnime(message string) error {
	fmt.Println("SendMessageUpdateAnime start....")
	err := config.SendMessage("UPDATE_KAFKA_MSG", message)
	if err != nil {
		return err
	}
	return nil
}
