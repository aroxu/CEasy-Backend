package server

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/aroxu/CEasy-Backend/db"
	"github.com/aroxu/CEasy-Backend/models"
	"github.com/aroxu/CEasy-Backend/utils"
)

//StartCrawl starts crawling to
func StartCrawl() {
	log.Print("Starting Crawling Code")
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
	mRmPrefix := regexp.MustCompile("\\[(.*?)\\]")

	str := strings.ReplaceAll(strings.ReplaceAll(data.Data.Content, "\r", ""), "\n", "")
	strPrefix := mRmPrefix.FindString(str)
	area := strings.Replace(strings.Replace(strPrefix, "[", "", -1), "]", "", -1)
	splitStr := strings.Split(str, "-송출지역-")
	areaDetail := "알수없음"
	if len(splitStr) > 1 {
		areaDetail = splitStr[1]
	}
	str = strings.Replace(splitStr[0], strPrefix, "", -1)

	timeParsed, err := time.Parse("2006-01-02 15:04:05", data.Data.Date)
	if err != nil {
		log.Print(err)
		return
	}

	idParsed, err := strconv.Atoi(data.Data.ID)
	if err != nil {
		log.Print(err)
		return
	}

	dbData := models.CeasyData{
		ID:         idParsed,
		Area:       strings.TrimSpace(area),
		AreaDetail: strings.TrimSpace(areaDetail),
		Content:    strings.TrimSpace(str),
		Date:       &timeParsed,
	}

	errInsertMsgToDb := db.InsertMsg(&dbData)
	if errInsertMsgToDb != nil {
		log.Fatalf("There was an error while insert message to Database Server. Check DB Setting and try again.")
	}
	log.Printf("End Parse ID: %d Area: %s\n", id, areaDetail)
}
