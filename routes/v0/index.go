package v0

import (
	c "github.com/B1ackAnge1/CEasy-Backend/controllers/v0"
	m "github.com/B1ackAnge1/CEasy-Backend/middlewares"
	"github.com/B1ackAnge1/CEasy-Backend/models/req"
	"github.com/gin-gonic/gin"
)

//InitRoutes initialize router for gin
func InitRoutes(g *gin.RouterGroup) {
	g.GET("/", m.VerifyQuery(&req.Search{}), c.Get)
	g.GET("/location", m.VerifyQuery(&req.Location{}), c.GetLocation)
}
