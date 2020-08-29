package res

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, data interface{}) {
	m := make(map[string]interface{})
	m["code"] = "SUCCESS"
	m["message"] = ""
	j, _ := json.Marshal(data)
	json.Unmarshal(j, &m)
	c.JSON(http.StatusOK, m)
}