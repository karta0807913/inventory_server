package router

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/karta0807913/inventory_server/model"
)

func ApiRouter(config RouterConfig) {
	db := config.DB
	router := config.Router

	router.GET("/item", func(c *gin.Context) {
		var item model.ItemTable
		search, err := item.GET(c, db)

		if err == nil {
			err = search.First(&item).Error
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "item not found",
				})
				return
			}
			c.JSON(200, item)
		} else {
			var max_limit = 20
			var limit int = max_limit
			slimit, ok := c.GetQuery("limit")
			if ok {
				limit, err = strconv.Atoi(slimit)
				if err != nil {
					limit = max_limit
				} else {
					if limit <= 0 || max_limit < limit {
						limit = max_limit
					}
				}
			}
			var item []model.ItemTable
			err = db.Limit(limit).Find(&item).Error
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "item not found",
				})
				return
			}
			c.JSON(200, item)
		}
	})

	router.PUT("/item", func(c *gin.Context) {
		var table model.ItemTable
		err := table.PUT(c, db)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
				"message": "update item error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{})
	})
}
