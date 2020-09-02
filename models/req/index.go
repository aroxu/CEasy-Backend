package req

type Location struct {
	Value string `form:"value" binding:"required"`
}

type Search struct {
	Area string `form:"area"`
	AreaDetail string `form:"area_detail"`
	Offset   int    `form:"offset"`
	Limit    int    `form:"limit"`
	Start string `form:"start"`
	End string `form:"end"`
}
