package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrType int

const (
	ERR_BAD_REQUEST ErrType = 1 + iota
	ERR_SERVER
	ERR_DUPLICATE
	ERR_AUTH
)

func SendError(c *gin.Context, errType ErrType, text string) {
	var Message, ErrCode string
	var Status int

	set := func(errCode, message string, status int) {
		ErrCode = errCode
		Message = message
		Status = status
	}

	switch errType {
	case ERR_BAD_REQUEST:
		set("ERR_BAD_REQUEST", text, http.StatusBadRequest)
	case ERR_SERVER:
		set("ERR_SERVER", text, http.StatusInternalServerError)
	case ERR_DUPLICATE:
		set("ERR_DUPLICATE", text, http.StatusConflict)
	case ERR_AUTH:
		set("ERR_AUTH", text, http.StatusUnauthorized)
	}

	c.JSON(Status, gin.H{
		"code":    ErrCode,
		"message": Message,
	})
	c.Abort()
}