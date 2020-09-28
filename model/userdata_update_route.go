package model

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// this file generate by go generate, please don't edit it
// data will put into struct
func (insert *UserData) Update(c *gin.Context, db *gorm.DB) error {
	type Body struct {
		Password string `json:"password" binding:"required"`

		Name *string `json:"nickname"`
	}
	var body Body
	err := c.ShouldBindJSON(&body)
	if err != nil {
		return err
	}
	insert.Password = body.Password

	selectField := make([]string, 0)

	if body.Name != nil {
		selectField = append(selectField, "name")
		insert.Name = *body.Name
	}

	if len(selectField) == 0 {
		return errors.New("rqeuire at least one option")
	}

	return db.Select(
		selectField[0], selectField[1:],
	).Where("password=?", body.Password).Updates(&insert).Error
}
