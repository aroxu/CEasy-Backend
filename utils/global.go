package utils

import (
	"github.com/B1ackAnge1/CEasy-Backend/models"
)

var (
	g_config *models.Config
)

func SetConfig(config *models.Config) {
	g_config = config
}

func GetConfig() *models.Config {
	return g_config
}
