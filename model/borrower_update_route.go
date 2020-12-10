package model

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// this file generate by go generate, please don't edit it
// data will put into struct
func (insert *Borrower) Update(c *gin.Context, db *gorm.DB) error {
	type Body struct {
		ID uint `json:"id" binding:"required"`

		Name  *string `json:"name"`
		Phone *string `json:"phone"`
	}
	var body Body
	err := c.ShouldBindJSON(&body)
	if err != nil {
		return err
	}
	insert.ID = body.ID

	selectField := make([]string, 0)

	if body.Name != nil {
		selectField = append(selectField, "name")
		insert.Name = *body.Name
	}

	if body.Phone != nil {
		selectField = append(selectField, "phone")
		insert.Phone = *body.Phone
	}

	if len(selectField) == 0 {
		return errors.New("require at least one option")
	}

	return db.Select(
		selectField[0], selectField[1:],
	).Where("borrowers.id=?", body.ID).Updates(&insert).Error
}
