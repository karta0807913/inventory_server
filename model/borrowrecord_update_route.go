package model

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// this file generate by go generate, please don't edit it
// data will put into struct
func (insert *BorrowRecord) Update(c *gin.Context, db *gorm.DB) error {
	type Body struct {
		ID uint `json:"id" binding:"required"`

		ReplyDate *time.Time `json:"reply_date"`
		Note      *string    `json:"note"`
		Returned  *bool      `json:"returned"`
	}
	var body Body
	err := c.ShouldBindJSON(&body)
	if err != nil {
		return err
	}
	insert.ID = body.ID

	selectField := make([]string, 0)

	if body.ReplyDate != nil {
		selectField = append(selectField, "reply_date")
		insert.ReplyDate = *body.ReplyDate
	}

	if body.Note != nil {
		selectField = append(selectField, "note")
		insert.Note = *body.Note
	}

	if body.Returned != nil {
		selectField = append(selectField, "returned")
		insert.Returned = *body.Returned
	}

	if len(selectField) == 0 {
		return errors.New("require at least one option")
	}

	return db.Select(
		selectField[0], selectField[1:],
	).Where("borrow_records.id=?", body.ID).Updates(&insert).Error
}
