package v0

import (
	"github.com/B1ackAnge1/CEasy-Backend/db"
	"github.com/B1ackAnge1/CEasy-Backend/models/req"
	resmodels "github.com/B1ackAnge1/CEasy-Backend/models/res"
	"github.com/B1ackAnge1/CEasy-Backend/utils/res"
	"github.com/gin-gonic/gin"
)

//GetLocation return location based on CBS Data
func GetLocation(c *gin.Context) {
	query := c.MustGet("query").(*req.Location)
	list, err := db.GetAreaMsg(query.Value)
	if err != nil{
		res.SendError(c, res.ErrServer, "ERROR")
	}
	res.Response(c, resmodels.Location{
		Data: list,
	})
}
