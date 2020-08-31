package utils

import (
	"github.com/B1ackAnge1/CEasy-Backend/models"
	"gorm.io/gorm"
)

var (
	gConfig *models.Config
	gDb     *gorm.DB
)

//SetConfig sets config from ceasy.config.json
func SetConfig(config *models.Config) {
	gConfig = config
}

//GetConfig gets config from ceasy.config.json
func GetConfig() *models.Config {
	return gConfig
}

//SetDB sets database config from ceasy.config.json
func SetDB(db *gorm.DB) {
	gDb = db
}

//GetDB gets database config from ceasy.config.json
func GetDB() *gorm.DB {
	return gDb
}
