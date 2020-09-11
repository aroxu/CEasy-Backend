package req

//Location CBS location structure
type Location struct {
	Value string `form:"value" binding:"required"`
}

//Search CBS filter and searching structure
type Search struct {
	Area       string `form:"area"`
	AreaDetail string `form:"area_detail"`
	Offset     int    `form:"offset"`
	Limit      int    `form:"limit"`
	Start      string `form:"start"`
	End        string `form:"end"`
}
