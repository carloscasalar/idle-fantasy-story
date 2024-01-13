package system

import (
	"github.com/gin-gonic/gin"
)

type Routes struct{}

func NewRoutes() *Routes {
	return &Routes{}
}

func (s *Routes) Register(g *gin.RouterGroup) {
	g.GET("/healthz", HealthHandler)
}
