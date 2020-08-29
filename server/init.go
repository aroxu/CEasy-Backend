package server

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/B1ackAnge1/CEasy-Backend/middlewares"
	"github.com/B1ackAnge1/CEasy-Backend/models"
	v0 "github.com/B1ackAnge1/CEasy-Backend/routes/v0"
	"github.com/B1ackAnge1/CEasy-Backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() {
	etcInit()
	applyConfig()
	initDB()
	startServer()
}

func etcInit() {
}

func applyConfig() {
	jsonFile, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	var config models.Config
	if err := json.Unmarshal(jsonFile, &config); err != nil {
		log.Fatal(err.Error())
	}
	log.Print("Successfully Opened config.json")
	utils.SetConfig(&config)
}

func startServer() {
	config := utils.GetConfig()
	if config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(middlewares.Cors())

	version0 := r.Group("/v0")
	v0.InitRoutes(version0)

	r.Run(":" + config.Port)

}

func initDB() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}
	utils.SetDB(db)
	log.Print("Successed To Connect Database")

	var models = []interface{}{&models.MsgData{}}
	db.AutoMigrate(models...)
	log.Print("Successfully performed AutoMigrate")
}
