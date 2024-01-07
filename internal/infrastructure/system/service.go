package system

import (
	"github.com/gin-gonic/gin"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) RegisterService(g *gin.RouterGroup) {
	g.GET("/healthz", HealthHandler)
}
