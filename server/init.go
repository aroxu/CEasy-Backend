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
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Init is actual entry point for this project.
func Init() {
	applyConfig()
	initDB()
	go testCode()
	startServer()
}

func testCode() {
}

func applyConfig() {
	jsonFile, err := ioutil.ReadFile("./ceasy.config.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	var config models.Config
	if err := json.Unmarshal(jsonFile, &config); err != nil {
		log.Fatal(err.Error())
	}
	log.Print("Successfully Opened ceasy.config.json")
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

	version0 := r.Group("/api/cbs/v0")
	v0.InitRoutes(version0)

	r.Run(":" + config.Port)

}

func initDB() {
	config := utils.GetConfig()
	dsn := config.DBUser + ":" + config.DBPass + "@tcp(" + config.DBHost + ":" + config.DBPort + ")/" + config.DBDatabase + "?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}
	utils.SetDB(db)
	log.Print("Successed To Connect Database")

	var models = []interface{}{&models.CeasyData{}}
	db.AutoMigrate(models...)
	log.Print("Successfully performed AutoMigrate")
}
