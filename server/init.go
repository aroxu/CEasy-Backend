package server

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/aroxu/CEasy-Backend/middlewares"
	"github.com/aroxu/CEasy-Backend/models"
	v0 "github.com/aroxu/CEasy-Backend/routes/v0"
	"github.com/aroxu/CEasy-Backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Init is actual entry point for this project.
func Init() {
	etcInit()
	applyConfig()
	initDB()
	startServer()
}

func etcInit() {
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

	errRunGin := r.Run(":" + config.Port)
	if errRunGin != nil {
		log.Fatalf("There was an error while running gin Server. Try to change config to DEBUG and check error.")
	}
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
	log.Print("Successfully Connected To Database")

	var ceasyDataModels = []interface{}{&models.CeasyData{}}
	errAutoMigrate := db.AutoMigrate(ceasyDataModels...)
	if errAutoMigrate != nil {
		log.Fatalf("There was an error while running Auto Migrate. Try to change config to DEBUG and check error.")
	}
	log.Print("Successfully performed AutoMigrate")
}
