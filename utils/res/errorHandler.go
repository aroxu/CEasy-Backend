package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrType int

const (
	ErrBadRequest ErrType = 1 + iota
	ErrServer
	ErrDuplicate
	ErrAuth
	ErrLimit
)

//SendError returns request error or other errors to request
func SendError(c *gin.Context, errType ErrType, text string) {
	var Message, ErrCode string
	var Status int

	set := func(errCode, message string, status int) {
		ErrCode = errCode
		Message = message
		Status = status
	}

	switch errType {
	case ErrBadRequest:
		set("ERR_BAD_REQUEST", text, http.StatusBadRequest)
	case ErrServer:
		set("ERR_SERVER", text, http.StatusInternalServerError)
	case ErrDuplicate:
		set("ERR_DUPLICATE", text, http.StatusConflict)
	case ErrAuth:
		set("ERR_AUTH", text, http.StatusUnauthorized)
	case ErrLimit:
		set("ERR_LIMIT", text, http.StatusBadRequest)
	}

	c.JSON(Status, gin.H{
		"code":    ErrCode,
		"message": Message,
	})
	c.Abort()
}
