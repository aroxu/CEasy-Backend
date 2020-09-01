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
	Result RtnResult `json:"rtnResult"`
	Data   CeasyData `json:"bbsMap"`
}

type RtnResult struct {
	Code string `json:"resultCode"`
	Msg  string `json:"resultMsg"`
}

type CeasyData struct {
	ID             int        `gorm:"primary_key;unique_index"`
	Title          string     `json:"sj" gorm:"unique_index"`
	Area           string     `json:"area" gorm:""`
	Content        string     `json:"cn"`
	IDStr          string     `json:"bbs_ordr" gorm:"-"`
	FirstRegist    *time.Time `json:"-"`
	FirstRegistStr string     `json:"frst_regist_dt"`
}
