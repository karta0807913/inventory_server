package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karta0807913/inventory_server/server"
	"gorm.io/gorm"
)

type RouterConfig struct {
	DB     *gorm.DB
	Router *gin.RouterGroup
}

func InitRouter(config RouterConfig) {
	db := config.DB
	router := config.Router
	commonRouter := router.Group("/")
	CommonRouter(RouterConfig{
		DB:     db,
		Router: commonRouter,
	})

	apiRouter := router.Group("/api", func(c *gin.Context) {
		session := c.MustGet("session").(server.Session)
		user_id := session.Get("user_id")
		if user_id == nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "please login",
			})
		}
	})
	ApiRouter(RouterConfig{
		DB:     db,
		Router: apiRouter,
	})
}
