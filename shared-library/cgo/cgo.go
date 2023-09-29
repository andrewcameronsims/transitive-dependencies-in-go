package cgo

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func HelloWorld() string {
	// This won't work at runtime unless you have Kafka running locally with this config,
	// but this code is just to demonstrate build time behaviour.
	c, _ := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	fmt.Println(c)

	return "Hello, world!"
}
