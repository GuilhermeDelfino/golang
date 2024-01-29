package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type Product struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func main() {

	messages := make(chan string)
	consumerQuantity := 3 // quantity of partitions
	for i := 1; i <= consumerQuantity; i++ {
		go consumer(messages, i)
	}

	for item := range messages {
		log.Println(item)
	}
}

func consumer(messages chan string, number int) {
	reader := kafka.NewReader(
		kafka.ReaderConfig{
			Brokers:     []string{"localhost:9092", "localhost:9093", "localhost:9094"},
			Topic:       "mytopic",
			GroupID:     "go-consumer",
			MinBytes:    10e3, // 10KB
			MaxBytes:    10e6, // 10MB
			StartOffset: kafka.FirstOffset,
		},
	)
	defer reader.Close()

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("Error reading message:", err)
		}

		// var product Product
		// json.Unmarshal(message.Value, &product)
		log.Println("Key", string(message.Key))
		log.Println("Value", string(message.Value))
		log.Println("Offset", message.Offset)
		item := fmt.Sprintf("Consumer %d\n", number)
		log.Print(item)
		log.Println("-----------")
		messages <- item
		// close(messages)
	}
}
