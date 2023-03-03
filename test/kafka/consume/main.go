package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {
	// Set up configuration
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Create consumer
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Println("Error creating consumer:", err)
		return
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			fmt.Println("Error closing consumer:", err)
		}
	}()

	// Subscribe to topic
	topic := "test-topic"
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		fmt.Println("Error subscribing to topic:", err)
		return
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			fmt.Println("Error closing partition consumer:", err)
		}
	}()

	// Consume messages
	for message := range partitionConsumer.Messages() {
		fmt.Printf("Received message with value %s\n", message.Value)
	}
}
