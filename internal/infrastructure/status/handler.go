package status

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "server is ready and healthy"})
}
