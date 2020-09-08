package model

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// this file generate by go generate, please don't edit it
// search options will put into struct
func (item *ItemTable) GET(c *gin.Context, db *gorm.DB) (*gorm.DB, error) {
	type Body struct {
		ItemID *string ``
	}
	var body Body
	err := c.ShouldBindQuery(&body)
	if err != nil {
		return nil, err
	}

	whereField := make([]string, 0)
	valueField := make([]interface{}, 0)

	if body.ItemID != nil {
		whereField = append(whereField, "ItemID=?")
		valueField = append(valueField, body.ItemID)
		item.ItemID = *body.ItemID
	}

	return db.Where(
		whereField,
	), nil
}
