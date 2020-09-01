package middlewares

import (
	"github.com/B1ackAnge1/CEasy-Backend/utils/res"
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

func VerifyQuery(data interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindQuery(data); err != nil {
			res.SendError(c, res.ErrBadRequest, err.Error())
			return
		}
		c.Set("query", data)
	}
}