package v0

import (
	"fmt"

	"github.com/B1ackAnge1/CEasy-Backend/models/req"
	resmodels "github.com/B1ackAnge1/CEasy-Backend/models/res"
	"github.com/B1ackAnge1/CEasy-Backend/utils/res"
	"github.com/gin-gonic/gin"
)

//GetLocation return location based on CBS Data
func GetLocation(c *gin.Context) {
	query := c.MustGet("query").(*req.Location)
	fmt.Println(query.Value)
	res.Response(c, resmodels.Empty{})
}
