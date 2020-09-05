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
		itemID, ok := c.GetQuery("item_id")
		if !ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "item_id missing",
			})
			return
		}
		item := model.ItemTable{}
		err := db.First(&item, "item_id=?", itemID).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "item not found",
			})
			return
		}
		c.JSON(http.StatusOK, item)
	})

	router.PUT("/item", func(c *gin.Context) {
		type Body struct {
			ItemID   string           `json:"item_id" binding:"required"`
			Location *string          `json:"location"`
			State    *model.ItemState `json:"state"`
		}
		var body Body
		err := c.ShouldBindJSON(&body)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "format error",
			})
			return
		}
		selectDB := db
		flag := false
		Select := make([]string, 0)
		if body.Location != nil {
			Select = append(Select, "Location")
			flag = true
		} else {
			body.Location = new(string)
		}
		if body.State != nil {
			Select = append(Select, "State")
			flag = true
		} else {
			body.State = &model.ItemState{}
		}
		if !flag {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "must contain one of location and state",
			})
			return
		}
		err = selectDB.Select(
			Select[0],
			Select[1:],
		).Where("item_id=?", body.ItemID).Updates(&model.ItemTable{
			State:    *body.State,
			Location: *body.Location,
		}).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
				"message": "update item error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{})
	})
}
