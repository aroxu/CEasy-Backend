package db

import (
	"github.com/B1ackAnge1/CEasy-Backend/models"
	"github.com/B1ackAnge1/CEasy-Backend/utils"
)

func Insert(data *models.MsgData) error {
	result := utils.GetDB().Create(&data)
	return Result.Error
}
