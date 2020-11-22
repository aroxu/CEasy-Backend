package db

import (
	"time"

	"github.com/aroxu/CEasy-Backend/models"
	"github.com/aroxu/CEasy-Backend/utils"
	"gorm.io/gorm"
)

//InsertMsg inserts data to database
func InsertMsg(data *models.CeasyData) error {
	result := utils.GetDB().Create(&data)
	return result.Error
}

//GetLast returns latest data from database
func GetLast() (int, error) {
	data := &models.CeasyData{}
	result := utils.GetDB().Order("id desc").Select("id").Find(&data)
	return data.ID, result.Error
}

//GetMsg returns CBS message by area from database.
func GetMsg(area, areaDetail, includes string, start, end *time.Time, limit, offset int) (*[]models.CeasyData, error) {
	var result *gorm.DB
	var data []models.CeasyData
	result = utils.GetDB().Limit(limit).Offset(offset).Order("id desc").Find(&data)
	if area != "" {
		result = result.Where("area LIKE ?", "%"+area+"%")
	}
	if areaDetail != "" {
		result = result.Where("area_detail LIKE ?", "%"+areaDetail+"%")
	}
	if includes != "" {
		result = result.Where("content LIKE ?", "%"+includes+"%")
	}
	if start != nil {
		result = result.Where("date >= ? ", start)
	}
	if end != nil {
		result = result.Where("date <= ? ", end)
	}
	result.Find(&data)
	return &data, result.Error
}

//GetMsgCount returns Total CBS message counts.
func GetMsgCount(area, areaDetail, includes string, start, end *time.Time, offset int) (int64, error) {
	var result *gorm.DB
	var count int64
	result = utils.GetDB().Table("ceasy_data")
	if area != "" {
		result = result.Where("area LIKE ?", "%"+area+"%")
	}
	if areaDetail != "" {
		result = result.Where("area_detail LIKE ?", "%"+areaDetail+"%")
	}
	if includes != "" {
		result = result.Where("content LIKE ?", "%"+includes+"%")
	}
	if start != nil {
		result = result.Where("date >= ? ", start)
	}
	if end != nil {
		result = result.Where("date <= ? ", end)
	}
	result.Count(&count)
	return count, result.Error
}

//GetAreaMsg returns all area's CBS message from database.
func GetAreaMsg(str string) ([]string, error) {
	data := make([]models.CeasyData, 0)
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
