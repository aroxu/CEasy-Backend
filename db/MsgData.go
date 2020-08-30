package db

import (
	"log"

	"github.com/B1ackAnge1/CEasy-Backend/models"
	"github.com/B1ackAnge1/CEasy-Backend/utils"
	"gorm.io/gorm"
)

//InsertMsg inserts data to database
func InsertMsg(data *models.MsgData) error {
	result := utils.GetDB().Create(&data)
	return result.Error
}

//GetLast returns latest data from database
func GetLast() (int, error) {
	data := &models.MsgData{}
	result := utils.GetDB().Order("id desc").Select("id").Find(&data)
	return data.ID, result.Error
}

//GetMsg returns CBS message by area from database.
func GetMsg(area string, limit, offset int) (*[]models.MsgData, error) {
	log.Print(limit)
	var result *gorm.DB
	data := []models.MsgData{}
	if area == "" {
		result = utils.GetDB().Limit(limit).Offset(offset).Order("id desc").Find(&data)
	} else {
		result = utils.GetDB().Where("area = ?", area).Limit(limit).Offset(offset).Order("id desc").Find(&data)
	}
	return &data, result.Error
}

//GetMsgCount returns Total CBS message counts.
func GetMsgCount(area string, offset int) (count int64, err error) {
	var result *gorm.DB
	if area == "" {
		result = utils.GetDB().Table("msg_data").Count(&count)
	} else {
		result = utils.GetDB().Table("msg_data").Where("area = ?", area).Count(&count)
	}
	err = result.Error
	return
}

//GetAreaMsg returns all area's CBS message from database.
func GetAreaMsg(str string) ([]string, error) {
	data := make([]models.MsgData, 0)
	result := utils.GetDB().Distinct("area").Order("area").Select("area").Where("area LIKE ?", "%"+str+"%").Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	areaList := make([]string, len(data))
	for i, d := range data {
		areaList[i] = d.Area
	}
	return areaList, nil
}
