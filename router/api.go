package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karta0807913/inventory_server/model"
)

func ApiRouter(config RouterConfig) {
	db := config.DB
	router := config.Router

	router.GET("/item", func(c *gin.Context) {
		var item model.ItemTable
		var result interface{}
		var err error
		_, ok := c.GetQuery("item_id")
		if ok {
			err = item.First(c, db)
			result = item
		} else {
			result, err = item.Find(c, db)
		}
		if err != nil {
			log.Printf("search item got error %s\n", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "item not found",
			})
			return
		}
		c.JSON(200, result)
	})

	router.PUT("/item", func(c *gin.Context) {
		var table model.ItemTable
		err := table.Update(c, db)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
				"message": "update item error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{})
	})
}
