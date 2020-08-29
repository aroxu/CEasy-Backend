package req

type Location struct {
	Value string `form:"value" binding:"required"`
}