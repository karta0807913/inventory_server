package model

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// this file generate by go generate, please don't edit it
// search options will put into struct
func (item *BorrowRecord) First(c *gin.Context, db *gorm.DB) error {
	type Body struct {
		ID         *uint `form:"id"`
		BorrowerID *uint `form:"borrower_id"`
	}

	var body Body
	err := c.ShouldBindQuery(&body)
	if err != nil {
		return err
	}

	whereField := make([]string, 0)
	valueField := make([]interface{}, 0)

	if body.ID != nil {
		whereField = append(whereField, "borrow_records.id=?")
		valueField = append(valueField, body.ID)
		item.ID = *body.ID
	}

	if body.BorrowerID != nil {
		whereField = append(whereField, "borrow_records.borrower_id=?")
		valueField = append(valueField, body.BorrowerID)
		item.BorrowerID = *body.BorrowerID
	}

	if len(valueField) < (0 + 0 + 1) {
		return errors.New("require option")
	}

	err = db.Where(
		strings.Join(whereField, " and "),
		valueField[0], valueField[1:],
	).First(item).Error
	return err
}
