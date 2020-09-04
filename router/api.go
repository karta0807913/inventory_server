package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karta0807913/inventory_server/model"
)

func ApiRouter(config RouterConfig) {
	db := config.DB
	router := config.Router

	router.GET("/item", func(c *gin.Context) {
		itemID, ok := c.GetQuery("item_id")
		if !ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "item_id missing",
			})
			return
		}
		item := model.ItemTable{}
		err := db.First(&item, "ItemID=?", itemID).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "item not found",
			})
		}
	})
}
