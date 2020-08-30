package v0

import (
	"github.com/B1ackAnge1/CEasy-Backend/db"

	"github.com/B1ackAnge1/CEasy-Backend/models/req"
	resmodels "github.com/B1ackAnge1/CEasy-Backend/models/res"
	"github.com/B1ackAnge1/CEasy-Backend/utils/res"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	query := c.MustGet("query").(*req.Search)
	if query.Limit == -1 {
		query.Limit = 10000000
	}
	if query.Limit == 0 {
		query.Limit = 20
	}
	data, err := db.GetMsg(query.Location, query.Limit, query.Offset)
	count, err2 := db.GetMsgCount(query.Location, query.Offset)
	if err != nil || err2 != nil {
		res.SendError(c, res.ERR_SERVER, "ERROR")
	}
	res.Response(c, resmodels.Search{
		Count: count,
		Data:  *data,
	})
}
