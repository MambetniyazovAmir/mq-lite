package main

import (
	"fmt"
	"time"

	"mq-lite/pkg/client"
)

func main() {
	c1, _ := client.Connect("localhost:8000")
	c2, _ := client.Connect("localhost:8000")

	msgCh, _ := c1.Subscribe("news")
	go func() {
		for msg := range msgCh {
			fmt.Println("→ Subscriber got:", msg)
		}
	}()

	time.Sleep(1 * time.Second)
	c2.Publish("news", "Hello from MQ-Lite SDK!")

	time.Sleep(2 * time.Second)
}
