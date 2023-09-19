package config

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
)

var (
	producer *kafka.Producer
)

func ConnectKafka() error {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
	kafkaBokers := os.Getenv("KAFKA_BROKERS")
	fmt.Println("[kafka] bokers:", kafkaBokers)
	fmt.Println("[kafka] Connecting to Kafka. Please wait a moment...")
	config := &kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BROKERS"),
		"client.id":         "go-kafka-producer",
	}

	var err error
	producer, err = kafka.NewProducer(config)
	if err != nil {
		return err
	}
	fmt.Println("[kafka] Connected to Kafka successfully!")
	return nil
}

// SendMessage sends a message to the specified Kafka topic.
func SendMessage(topic string, message string) error {
	kafkaMessage := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}

	err := producer.Produce(kafkaMessage, nil)
	if err != nil {
		return err
	}

	return nil
}

// CloseKafka closes the Kafka producer.
func CloseKafka() {
	if producer != nil {
		producer.Close()
	}
}
