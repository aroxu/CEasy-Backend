package req

type Location struct {
	Location string `form:"location" binding:"required"`
}