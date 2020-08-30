package server

import (
	"fmt"
	"log"

	"github.com/B1ackAnge1/CEasy-Backend/db"
	"github.com/B1ackAnge1/CEasy-Backend/utils"
)

//StartCrawl starts crawling to
func StartCrawl() {
	id, err := utils.GetLastMsgID()
	if err != nil {
		log.Print(err)
		return
	}
	lastID, err := db.GetLast()
	if err != nil {
		log.Print(err)
		lastID = 0
	}
	fmt.Println("Start id: ", id, " Last ID: ", lastID)
	for ; id > lastID; id-- {
		fmt.Println("Start Parse ID: ", id)
		data, err := utils.GetDetailMsg(id)
		if err != nil {
			log.Print(err)
			return
		}
		go func() {
			data.Data.Area = utils.ParseStringInBetween(data.Data.Content, "[", "]")
			db.InsertMsg(&data.Data)
			fmt.Printf("End Parse ID: %d Area: %s\n", id, data.Data.Area)
		}()
	}
	log.Print("End Crawl")
}
