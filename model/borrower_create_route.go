package model

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// this file generate by go generate, please don't edit it
// data will put into struct
func (insert *Borrower) Create(c *gin.Context, db *gorm.DB) error {
	type Body struct {
		Name  string `json:"name" binding:"required"`
		Phone string `json:"phone" binding:"required"`
	}
	var body Body
	err := c.ShouldBindJSON(&body)
	if err != nil {
		return err
	}

	selectField := []string{
		"name",
		"phone",
	}

	insert.Name = body.Name
	insert.Phone = body.Phone

	return db.Select(
		selectField[0], selectField[1:],
	).Create(&insert).Error
}
