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

	err = db.Where(
		strings.Join(whereField, "and"),
		valueField,
	).First(item).Error
	return err
}
