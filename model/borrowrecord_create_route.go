package model

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// this file generate by go generate, please don't edit it
// data will put into struct
func (insert *BorrowRecord) Create(c *gin.Context, db *gorm.DB) error {
	type Body struct {
		BorrowDate time.Time `json:"borrow_date" binding:"required"`
		ReplyDate  time.Time `json:"reply_date" binding:"required"`

		Borrower   *Borrower `json:"borrower"`
		BorrowerID *uint     `json:"borrower_id"`
		Note       *string   `json:"note"`
	}
	var body Body
	err := c.ShouldBindJSON(&body)
	if err != nil {
		return err
	}

	selectField := []string{
		"borrow_date",
		"reply_date",
	}

	if body.Borrower != nil {
		selectField = append(selectField, "borrower")
		insert.Borrower = *body.Borrower
	}

	if body.BorrowerID != nil {
		selectField = append(selectField, "borrower_id")
		insert.BorrowerID = *body.BorrowerID
	}

	if body.Note != nil {
		selectField = append(selectField, "note")
		insert.Note = *body.Note
	}

	if len(selectField) == 2 {
		return errors.New("rqeuire at least one option")
	}

	insert.BorrowDate = body.BorrowDate
	insert.ReplyDate = body.ReplyDate

	return db.Select(
		selectField[0], selectField[1:],
	).Create(&insert).Error
}
