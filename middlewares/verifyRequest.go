package middlewares

import (
	"github.com/aroxu/CEasy-Backend/utils/res"
	"github.com/gin-gonic/gin"
)

//func VerifyBody(data interface{}) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		if err := c.ShouldBindJSON(data); err != nil {
//			res.SendError(c, res.ErrBadRequest, err.Error())
//			return
//		}
//		c.Set("body", data)
//	}
//}

//VerifyQuery filters query data based on request header
func VerifyQuery(data interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindQuery(data); err != nil {
			res.SendError(c, res.ErrBadRequest, err.Error())
			return
		}
		c.Set("query", data)
	}
}
