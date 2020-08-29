package models

import (
	"time"
)

type SelectBbsList struct {
	Result RtnResult `json:"rtnResult"`
	Data   []struct {
		ID       int       `json:"BBS_ORDR"`
		CreateAt time.Time `json:"FRST_REGIST_DT"`
		Title    string    `json:"SJ"`
	} `json:"bbsList"`
}

type SelectBbsView struct {
	Result RtnResult `json:"rtnResult"`
	Data   MsgData   `json:"bbsMap"`
}

type RtnResult struct {
	Code string `json:"resultCode"`
	Msg  string `json:"resultMsg"`
}

type MsgData struct {
	ID       int       `json:"bbs_ordr"`
	Title    string    `json:"sj" gorm:"unique_index"`
	Content  string    `json:"cn"`
	CreateAt time.Time `json:"frst_regist_dt"`
}
