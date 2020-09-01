package server

import (
	"log"

	"github.com/B1ackAnge1/CEasy-Backend/db"
	"github.com/B1ackAnge1/CEasy-Backend/models"
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
	log.Print("Start id: ", id, " Last ID: ", lastID)
	for ; id > lastID; id-- {
		log.Print("Start Parse ID: ", id)
		data, err := utils.GetDetailMsg(id)
		if err != nil {
			log.Print(err)
			return
		}
		go crawlInsertDB(id, data)
	}
	log.Print("End Crawl")
}

func crawlInsertDB(id int, data *models.SelectBbsView) {
	data.Data.Area = utils.ParseStringInBetween(data.Data.Content, "[", "]")
	errInsertMsgToDb := db.InsertMsg(&data.Data)
	if errInsertMsgToDb != nil {
		log.Fatalf("There was an error while insert message to Database Server. Check DB Setting and try again.")
	}
	log.Printf("End Parse ID: %d Area: %s\n", id, data.Data.Area)
}
