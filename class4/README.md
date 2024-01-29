# Class 4 - Kafka - Consume and produce

Kafka-go library
`go get github.com/segmentio/kafka-go`

### Producer

```
brokerAddress := "localhost:9092"
topic := "my-topic"

// Create a Kafka producer
writer := kafka.NewWriter(kafka.WriterConfig{
    Brokers:  []string{brokerAddress},
    Topic:    topic,
    Balancer: &kafka.LeastBytes{},
})

defer writer.Close()

// Produce a message to the Kafka topic
message := kafka.Message{
    Key:   []byte("key"),
    Value: []byte("Hello, Kafka from Golang!"),
}

err := writer.WriteMessages(context.Background(), message)
if err != nil {
    log.Fatal("Error writing message:", err)
}

fmt.Println("Message sent successfully!")
```

### Consumer

```
brokerAddress := "localhost:9092"
topic := "my-topic"
groupID := "my-group"

// Create a Kafka consumer
reader := kafka.NewReader(kafka.ReaderConfig{
    Brokers:  []string{brokerAddress},
    Topic:    topic,
    GroupID:  groupID,
    MinBytes: 10e3, // 10KB
    MaxBytes: 10e6, // 10MB
})

defer reader.Close()

// Continuously consume messages from the Kafka topic
for {
    message, err := reader.ReadMessage(context.Background())
    if err != nil {
        log.Fatal("Error reading message:", err)
    }

    fmt.Printf("Received message: %s\n", message.Value)
}
```
