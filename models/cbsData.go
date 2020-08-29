package models

//CBSData is structure for CBS Raw Data
type CBSData struct {
	Number       int    `json:"RNUM"`
	ModifiedDate string `json:"LAST_MODF_DT"`
	RegisterDate string `json:"FRST_REGIST_DT"`
	SJ           string `json:"SJ"`
	IndexNumber  int    `json:"IDX_NO"`
	Content      string `json:"CONT"`
	NM           string `json:"USR_NM"`
	ExpsrAt      string `json:"USR_EXPSR_AT"`
}
