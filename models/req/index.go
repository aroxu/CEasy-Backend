package req

type Location struct {
	Value string `form:"value" binding:"required"`
}

type Search struct {
	Location string `form:"location"`
	Offset   int    `form:"offset"`
	Limit    int    `form:"limit"`
}
