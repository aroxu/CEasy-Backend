package main

import (
	"time"

	"github.com/aroxu/CEasy-Backend/server"
)

func main() {
	go func() {
		for {
			server.StartCrawl()
			time.Sleep(time.Minute * 1)
		}
	}()
	server.Init()
}
