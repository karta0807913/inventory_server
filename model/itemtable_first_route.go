package model

import (
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// this file generate by go generate, please don't edit it
// search options will put into struct
func (item *ItemTable) First(c *gin.Context, db *gorm.DB) error {
	type Body struct {
		ItemID string `form:"item_id" binding:"required"`

		Name *string `form:"name"`
	}

	var body Body
	err := c.ShouldBindQuery(&body)
	if err != nil {
		return err
	}

	whereField := []string{
		"item_tables.item_id=?",
	}
	valueField := []interface{}{
		body.ItemID,
	}

	item.ItemID = body.ItemID

	if body.Name != nil {
		whereField = append(whereField, "item_tables.name=?")
		valueField = append(valueField, body.Name)
		item.Name = *body.Name
	}

	err = db.Where(
		strings.Join(whereField, " and "),
		valueField[0], valueField[1:],
	).First(item).Error
	return err
}
