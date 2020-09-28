package model

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// this file generate by go generate, please don't edit it
// search options will put into struct
func (item *ItemTable) GET(c *gin.Context, db *gorm.DB) (*gorm.DB, error) {
	type Body struct {
		ItemID *string    `form:"item_id"`
		State  *ItemState `form:"state"`
	}
	var body Body
	err := c.ShouldBindQuery(&body)
	if err != nil {
		return nil, err
	}

	whereField := make([]string, 0)
	valueField := make([]interface{}, 0)

	if body.ItemID != nil {
		whereField = append(whereField, "item_id=?")
		valueField = append(valueField, body.ItemID)
		item.ItemID = *body.ItemID
	}

	if body.State != nil {
		whereField = append(whereField, "state=?")
		valueField = append(valueField, body.State)
		item.State = *body.State
	}

	if len(valueField) == 0 {
		return nil, errors.New("must have one options")
	}

	return db.Where(
		strings.Join(whereField, "and"),
		valueField,
	), nil
}
