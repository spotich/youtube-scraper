package main

import (
	"fmt"
	_ "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("hello from consumer")
}
