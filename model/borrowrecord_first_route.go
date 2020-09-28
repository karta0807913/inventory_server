package model

import (
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// this file generate by go generate, please don't edit it
// search options will put into struct
func (item *BorrowRecord) First(c *gin.Context, db *gorm.DB) error {
	type Body struct {
		ID        *uint      `form:"id"`
		ReplyDate *time.Time `form:"reply_date"`
	}

	var body Body
	err := c.ShouldBindQuery(&body)
	if err != nil {
		return err
	}

	whereField := make([]string, 0)
	valueField := make([]interface{}, 0)

	if body.ID != nil {
		whereField = append(whereField, "id=?")
		valueField = append(valueField, body.ID)
		item.ID = *body.ID
	}

	if body.ReplyDate != nil {
		whereField = append(whereField, "reply_date=?")
		valueField = append(valueField, body.ReplyDate)
		item.ReplyDate = *body.ReplyDate
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
