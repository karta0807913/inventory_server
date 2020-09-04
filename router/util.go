package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func bodyMissing(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": "request body not correct",
	})
}
