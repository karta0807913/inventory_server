package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterConfig struct {
	DB     *gorm.DB
	Router *gin.RouterGroup
}

func InitRouter(config RouterConfig) {
	db := config.DB
	router := config.Router
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "http://172.17.0.2:3000")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
	})
	commonRouter := router.Group("/")
	CommonRouter(RouterConfig{
		DB:     db,
		Router: commonRouter,
	})

	apiRouter := router.Group("/api", func(c *gin.Context) {
		// session := c.MustGet("session").(serverutil.Session)
		// user_id := session.Get("user_id")
		// if user_id == nil {
		// 	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		// 		"message": "please login",
		// 	})
		// }
	})
	ApiRouter(RouterConfig{
		DB:     db,
		Router: apiRouter,
	})
}
