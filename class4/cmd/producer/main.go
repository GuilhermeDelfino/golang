package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"
)

type Product struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func main() {

	writer := kafka.NewWriter(
		kafka.WriterConfig{
			Brokers:          []string{"localhost:9092", "localhost:9093", "localhost:9094"},
			Topic:            "mytopic",
			Balancer:         &kafka.LeastBytes{},
			CompressionCodec: &compress.SnappyCodec,
			RequiredAcks:     1,
		},
	)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Type the product name: (Crtl + C to exit)")

		productName, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		randomId := rand.Int()

		product := Product{ID: randomId, Title: productName}

		json, err := json.Marshal(product)
		if err != nil {
			panic(err)
		}
		message := kafka.Message{
			Key:   []byte(strconv.Itoa(product.ID)),
			Value: json,
		}

		err = writer.WriteMessages(context.Background(), message)

		if err != nil {
			log.Fatal("Error writing message: ", err)
		}

		log.Println("Message sent")
		fmt.Println("-------------")
	}
	log.Println("Bye!")
}
