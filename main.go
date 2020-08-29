package main

import (
	"time"

	"github.com/B1ackAnge1/CEasy-Backend/server"
	"github.com/B1ackAnge1/CEasy-Backend/utils"
)

func main() {
	go func() {
		for {
			time.Sleep(time.Minute * 3)
			utils.StartCrawl()
		}
	}()
	server.Init()
}
