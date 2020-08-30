package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/B1ackAnge1/CEasy-Backend/models"
	json "github.com/json-iterator/go"
)

//GetDetailMsg get id (integet) and returns map with interfaced bbs searchinfo.
func GetDetailMsg(id int) (*models.SelectBbsView, error) {
	reqbody := &map[interface{}]interface{}{
		"bbs_searchInfo": map[interface{}]interface{}{
			"bbs_no":   "63",
			"bbs_ordr": id,
		},
	}
	reqbytes, err := json.Marshal(reqbody)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", "https://www.safekorea.go.kr/idsiSFK/bbs/user/selectBbsView.do", bytes.NewBuffer(reqbytes))
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	data := &models.SelectBbsView{}
	if err = json.Unmarshal(respBody, &data); err != nil {
		return nil, err
	}
	id, errStrconv := strconv.Atoi(data.Data.IDStr)
	if errStrconv != nil {
		return nil, errStrconv
	}
	data.Data.ID = id
	return data, nil
}

//GetLastMsgID return integar type of Latest bbs id and error(nil | error)
func GetLastMsgID() (int, error) {
	reqbody := &map[interface{}]interface{}{
		"bbs_searchInfo": map[interface{}]interface{}{
			"bbs_no":   "63",
			"pageUnit": "1",
		},
	}
	reqbytes, err := json.Marshal(reqbody)
	if err != nil {
		return 0, err
	}
	req, err := http.NewRequest("POST", "http://www.safekorea.go.kr/idsiSFK/bbs/user/selectBbsList.do", bytes.NewBuffer(reqbytes))
	if err != nil {
		return 0, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	data := &models.SelectBbsList{}
	if err = json.Unmarshal(respBody, &data); err != nil {
		return 0, err
	}
	return data.Data[0].ID, nil
}
