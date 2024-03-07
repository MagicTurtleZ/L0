package main

import (
	"fmt"
	"os"

	"github.com/nats-io/stan.go"
)

func main() {
	sh, err := stan.Connect("test-cluster", "test-pub")
	if err != nil {
		fmt.Println("the publisher cannot connect")
	}
	data, err := os.ReadFile("materials\\model.json")

	if err != nil {
		fmt.Println("canno`t read file")
	}

	sh.Publish("L0_enjoyer", data)
	fmt.Println("Сообщение отправлено")
	sh.Close()
}

