package main

import (
	"time"

	"github.com/B1ackAnge1/CEasy-Backend/server"
)

func main() {
	go func() {
		for {
			server.StartCrawl()
			time.Sleep(time.Minute * 3)
		}
	}()
	server.Init()
}
