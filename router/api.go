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
		_, itemOk := c.GetQuery("item_id")
		_, idOk := c.GetQuery("id")
		if itemOk || idOk {
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

	router.POST("/borrower", func(c *gin.Context) {
		var borrower model.Borrower
		err := borrower.Create(c, db)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
				"message": "create borrower failed",
			})
			return
		}
		c.JSON(http.StatusOK, borrower)
	})

	router.GET("/borrower", func(c *gin.Context) {
		var result interface{}
		var err error
		var borrower model.Borrower
		_, ok := c.GetQuery("id")
		if ok {
			err = borrower.First(c, db)
			result = borrower
		} else {
			result, err = borrower.Find(c, db)
		}
		if err != nil {
			log.Printf("search borrower got error %s\n", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "borrower not found",
			})
			return
		}
		c.JSON(http.StatusOK, result)
	})

	router.PUT("/borrower", func(c *gin.Context) {
		var borrower model.Borrower
		err := borrower.Update(c, db)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "update error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{})
	})

	router.GET("/borrow_record", func(c *gin.Context) {
		var result interface{}
		var err error
		var borrower model.BorrowRecord
		_, ok := c.GetQuery("id")
		if ok {
			err = borrower.First(c, db)
			result = borrower
		} else {
			result, err = borrower.Find(c, db)
		}
		if err != nil {
			log.Printf("search got error %s\n", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "borrow record not found",
			})
			return
		}
		c.JSON(http.StatusOK, result)
	})

	router.PUT("/borrow_record", func(c *gin.Context) {
		var borrower model.BorrowRecord
		err := borrower.Update(c, db)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "update error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{})
	})

	router.POST("/borrow_record", func(c *gin.Context) {
		var borrower model.BorrowRecord
		err := borrower.Create(c, db)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "create borrower failed",
			})
			return
		}
		c.JSON(http.StatusOK, borrower)
	})

	router.GET("/borrower_fuzzy", func(c *gin.Context) {
		var result []model.Borrower
		name, ok := c.GetQuery("name")
		query := db.Limit(20)
		if ok {
			query = query.Where("name like ?", name+"%")
		}
		phone, ok := c.GetQuery("phone")
		if ok {
			query = query.Where("phone like ?", phone+"%")
		}
		err := query.Find(&result).Error
		if err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusOK, result)
	})
}
