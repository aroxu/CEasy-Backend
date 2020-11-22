package v0

import (
	"github.com/aroxu/CEasy-Backend/db"
	"github.com/aroxu/CEasy-Backend/models/req"
	resmodels "github.com/aroxu/CEasy-Backend/models/res"
	"github.com/aroxu/CEasy-Backend/utils/res"
	"github.com/gin-gonic/gin"
)

//GetLocation return location based on CBS Data
func GetLocation(c *gin.Context) {
	query := c.MustGet("query").(*req.Location)
	list, err := db.GetAreaMsg(query.Value)
	if err != nil {
		res.SendError(c, res.ErrServer, "ERROR")
		return
	}
	res.Response(c, resmodels.Location{
		Data: list,
	})
}
