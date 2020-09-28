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
		BorrowerID uint      `json:"borrower_id" binding:"required"`
		BorrowDate time.Time `json:"borrow_date" binding:"required"`
		ReplyDate  time.Time `json:"reply_date" binding:"required"`
		Note       string    `json:"note" binding:"required"`

		Borrower *Borrower `json:"borrower"`
	}
	var body Body
	err := c.ShouldBindJSON(&body)
	if err != nil {
		return err
	}

	selectField := []string{
		"borrower_id",
		"borrow_date",
		"reply_date",
		"note",
	}

	if body.Borrower != nil {
		selectField = append(selectField, "borrower")
		insert.Borrower = *body.Borrower
	}

	if len(selectField) == 4 {
		return errors.New("rqeuire at least one option")
	}

	insert.BorrowerID = body.BorrowerID
	insert.BorrowDate = body.BorrowDate
	insert.ReplyDate = body.ReplyDate
	insert.Note = body.Note

	return db.Select(
		selectField[0], selectField[1:],
	).Create(&insert).Error
}
