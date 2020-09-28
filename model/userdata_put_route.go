package model

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// this file generate by go generate, please don't edit it
// data will put into struct
func (insert *UserData) PUT(c *gin.Context, db *gorm.DB) error {
	type Body struct {
		Password string `json:"password" binding:"required"`

		Name *string `json:"nickname"`
	}
	var body Body
	err := c.ShouldBindJSON(&body)
	if err != nil {
		return err
	}

	selectField := make([]string, 0)

	if body.Name == nil {
		body.Name = new(string)
	} else {
		selectField = append(selectField, "Name")
	}

	if len(selectField) == 0 {
		return errors.New("rqeuire at least one option")
	}

	insert.Password = body.Password
	insert.Name = *body.Name

	return db.Select(
		selectField[0], selectField[1:],
	).Where("password=?", body.Password).Updates(&insert).Error
}
