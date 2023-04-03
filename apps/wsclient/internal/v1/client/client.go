package client

import (
	"fmt"

	"github.com/atharva29/go-template/pkg/kafka"
)

func PrintClient() {
	kafka.KafkaInitialize()
	fmt.Println("I am client")
}
