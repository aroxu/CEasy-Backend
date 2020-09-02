package models

import "time"

type SelectBbsList struct {
	Result RtnResult `json:"rtnResult"`
	Data   []struct {
		ID       int    `json:"BBS_ORDR"`
		CreateAt string `json:"FRST_REGIST_DT"`
		Title    string `json:"SJ"`
	} `json:"bbsList"`
}

type SelectBbsView struct {
	Result RtnResult         `json:"rtnResult"`
	Data   CeasyDataForParse `json:"bbsMap"`
}

type RtnResult struct {
	Code string `json:"resultCode"`
	Msg  string `json:"resultMsg"`
}

type CeasyData struct {
	ID         int    `gorm:"primary_key;unique_index"`
	Area       string `json:"area" gorm:""`
	AreaDetail string `json:"area_detail" gorm:""`
	Content    string `json:"content"`
	Date       *time.Time
}

type CeasyDataForParse struct {
	ID      string `json:"bbs_ordr"`
	Content string `json:"cn"`
	Date    string `json:"frst_regist_dt"`
}
