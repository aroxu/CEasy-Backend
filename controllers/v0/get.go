package v0

import (
	"time"

	"github.com/aroxu/CEasy-Backend/db"
	"github.com/aroxu/CEasy-Backend/utils"

	"github.com/aroxu/CEasy-Backend/models/req"
	resmodels "github.com/aroxu/CEasy-Backend/models/res"
	"github.com/aroxu/CEasy-Backend/utils/res"
	"github.com/gin-gonic/gin"
)

//Get returns query from database.
func Get(c *gin.Context) {
	query := c.MustGet("query").(*req.Search)
	if c.Request.URL.Query().Get("limit") == "" {
		query.Limit = utils.GetConfig().DefaultLimit
	}
	if c.Request.URL.Query().Get("area") == "" {
		query.Area = ""
	}
	if c.Request.URL.Query().Get("area_detail") == "" {
		query.AreaDetail = ""
	}
	if c.Request.URL.Query().Get("offset") == "" {
		query.Offset = 0
	}
	if c.Request.URL.Query().Get("start") == "" {
		query.Start = ""
	}
	if c.Request.URL.Query().Get("end") == "" {
		query.End = ""
	}
	var start, end *time.Time
	start, end = nil, nil
	if query.Start != "" {
		parsed, errTime := time.Parse("2006-01-02 15:04:05", query.Start)
		if errTime != nil {
			res.SendError(c, res.ErrBadRequest, "시간을 yyyy-mm-dd hh:mm:ss 형식으로 적어주세요.")
			return
		}
		start = &parsed
	}
	if query.End != "" {
		parsed, errTime := time.Parse("2006-01-02 15:04:05", query.End)
		if errTime != nil {
			res.SendError(c, res.ErrBadRequest, "시간을 yyyy-mm-dd hh:mm:ss 형식으로 적어주세요.")
			return
		}
		end = &parsed
	}

	var requestCount = query.Limit
	if query.Limit == 0 {
		requestCount = utils.GetConfig().DefaultLimit
	}

	data, err := db.GetMsg(query.Area, query.AreaDetail, start, end, requestCount, query.Offset)
	count, err2 := db.GetMsgCount(query.Area, query.AreaDetail, start, end, query.Offset)
	if err != nil || err2 != nil {
		res.SendError(c, res.ErrServer, "There was a problem with Database")
		return
	}
	res.Response(c, resmodels.Search{
		Count: count,
		Data:  *data,
	})
}
