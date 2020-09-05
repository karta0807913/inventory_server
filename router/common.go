package router

import (
	"crypto/sha256"
	"encoding/base64"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karta0807913/inventory_server/model"
	"github.com/karta0807913/inventory_server/server"
)

func saltPassword(str string) string {
	encoder := sha256.New()
	password := base64.StdEncoding.EncodeToString(
		encoder.Sum(
			[]byte(str),
		),
	)
	return password
}

func CommonRouter(config RouterConfig) {
	router := config.Router
	db := config.DB
	router.POST("/login", func(c *gin.Context) {
		var body model.UserData
		err := c.ShouldBindJSON(&body)
		if err != nil {
			log.Println(err)
			bodyMissing(c)
			return
		}
		if body.Account == "" || body.Password == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "account or password missing",
			})
			return
		}
		err = db.Select(
			"ID", "Account", "Name",
		).First(
			&body, "Account=? and Password=?",
			body.Account, saltPassword(body.Password),
		).Error
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "account exists or too long",
			})
			return
		}
		session := c.MustGet("session").(server.Session)
		session.Set("user_id", body.ID)
		c.JSON(http.StatusOK, gin.H{
			"Name": body.Name,
		})
	})

	router.POST("/sign_up", func(c *gin.Context) {
		var body model.UserData
		err := c.ShouldBindJSON(&body)
		if err != nil {
			bodyMissing(c)
			return
		}
		if body.Account == "" || body.Password == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "account or password missing",
			})
			return
		}
		if body.Name == "" {
			body.Name = "No name"
		}
		body.Password = saltPassword(body.Password)
		err = db.Select("Account", "Password", "Name").Create(&body).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{})
	})
}