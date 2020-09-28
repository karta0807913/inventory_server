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

		BorrowerID *uint      `json:"borrower_id"`
		BorrowDate *time.Time `json:"borrow_date"`
		ReplyDate  *time.Time `json:"reply_date"`
		Note       *string    `json:"note"`
	}
	var body Body
	err := c.ShouldBindJSON(&body)
	if err != nil {
		return err
	}
	insert.ID = body.ID

	selectField := make([]string, 0)

	if body.BorrowerID != nil {
		selectField = append(selectField, "borrower_id")
		insert.BorrowerID = *body.BorrowerID
	}

	if body.BorrowDate != nil {
		selectField = append(selectField, "borrow_date")
		insert.BorrowDate = *body.BorrowDate
	}

	if body.ReplyDate != nil {
		selectField = append(selectField, "reply_date")
		insert.ReplyDate = *body.ReplyDate
	}

	if body.Note != nil {
		selectField = append(selectField, "note")
		insert.Note = *body.Note
	}

	if len(selectField) == 0 {
		return errors.New("rqeuire at least one option")
	}

	return db.Select(
		selectField[0], selectField[1:],
	).Where("id=?", body.ID).Updates(&insert).Error
}
