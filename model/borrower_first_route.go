package model

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// this file generate by go generate, please don't edit it
// search options will put into struct
func (item *Borrower) First(c *gin.Context, db *gorm.DB) error {
	type Body struct {
		ID    *uint   `form:"id"`
		Name  *string `form:"name"`
		Phone *string `form:"phone"`
	}

	var body Body
	err := c.ShouldBindQuery(&body)
	if err != nil {
		return err
	}

	whereField := make([]string, 0)
	valueField := make([]interface{}, 0)

	if body.ID != nil {
		whereField = append(whereField, "borrowers.id=?")
		valueField = append(valueField, body.ID)
		item.ID = *body.ID
	}

	if body.Name != nil {
		whereField = append(whereField, "borrowers.name=?")
		valueField = append(valueField, body.Name)
		item.Name = *body.Name
	}

	if body.Phone != nil {
		whereField = append(whereField, "borrowers.phone=?")
		valueField = append(valueField, body.Phone)
		item.Phone = *body.Phone
	}

	if len(valueField) == 0 {
		return errors.New("require at least one option")
	}

	err = db.Where(
		strings.Join(whereField, "and"),
		valueField,
	).First(item).Error
	return err
}
