package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "server is ready and healthy"})
}
