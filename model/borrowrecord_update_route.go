package model

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (insert *BorrowRecord) Update(c *gin.Context, db *gorm.DB) error {
	type Body struct {
		ID uint `json:"id" binding:"required"`

		BorrowerID *uint      `json:"borrower_id"`
		ReplyDate  *time.Time `json:"reply_date"`
		Note       *string    `json:"note"`
		Returned   *bool      `json:"returned"`
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

	if body.Note != nil {
		selectField = append(selectField, "note")
		insert.Note = *body.Note
	}

	if body.Returned != nil {
		selectField = append(selectField, "reply_date")
		if *body.Returned {
			t := time.Now()
			insert.ReplyDate = &t
		} else {
			insert.ReplyDate = nil
		}
	}

	if body.ReplyDate != nil {
		selectField = append(selectField, "reply_date")
		insert.ReplyDate = body.ReplyDate
	}

	if len(selectField) < (0 + 0 + 1) {
		return errors.New("require at least one option")
	}

	return db.Select(
		selectField[0], selectField[1:],
	).Where("borrow_records.id=?", body.ID).Updates(&insert).Error
}
