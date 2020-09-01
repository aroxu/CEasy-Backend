package res

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, data interface{}) {
	m := make(map[string]interface{})
	m["code"] = "SUCCESS"
	m["message"] = ""
	j, _ := json.Marshal(data)
	errJsonUnmarshal := json.Unmarshal(j, &m)
	if errJsonUnmarshal != nil {
		log.Fatalf("Failed during Unmarshaling json")
	}
	c.JSON(http.StatusOK, m)
}