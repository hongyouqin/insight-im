package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {
	// Set up configuration
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	// Create producer
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Println("Error creating producer:", err)
		return
	}
	defer func() {
		if err := producer.Close(); err != nil {
			fmt.Println("Error closing producer:", err)
		}
	}()

	// Send message
	message := &sarama.ProducerMessage{
		Topic: "test-topic",
		Value: sarama.StringEncoder("hello, kafka!"),
	}
	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}
	fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
}
