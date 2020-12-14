package model

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// this file generate by go generate, please don't edit it
// data will put into struct
func (insert *BorrowRecord) Create(c *gin.Context, db *gorm.DB) error {
	type Body struct {
		BorrowerID uint      `json:"borrower_id" binding:"required"`
		ItemID     uint      `json:"item_id" binding:"required"`
		BorrowDate time.Time `json:"borrow_date" binding:"required"`

		ReplyDate *time.Time `json:"reply_date"`
		Note      *string    `json:"note"`
	}
	var body Body
	err := c.ShouldBindJSON(&body)
	if err != nil {
		return err
	}

	selectField := []string{
		"borrower_id",
		"item_id",
		"borrow_date",
	}

	if body.ReplyDate != nil {
		selectField = append(selectField, "reply_date")
		insert.ReplyDate = *body.ReplyDate
	}

	if body.Note != nil {
		selectField = append(selectField, "note")
		insert.Note = *body.Note
	}

	insert.BorrowerID = body.BorrowerID
	insert.ItemID = body.ItemID
	insert.BorrowDate = body.BorrowDate

	return db.Select(
		selectField[0], selectField[1:],
	).Create(&insert).Error
}
