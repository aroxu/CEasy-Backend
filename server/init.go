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
)

func Init() {
	etcInit()
	applyConfig()
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
