package model

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// this file generate by go generate, please don't edit it
// data will put into struct
func (insert *ItemTable) Update(c *gin.Context, db *gorm.DB) error {
	type Body struct {
		ItemID string `json:"item_id" binding:"required"`

		AgeLimit *uint      `json:"age_limit"`
		Location *string    `json:"location"`
		State    *ItemState `json:"state"`
		Note     *string    `json:"note"`
	}
	var body Body
	err := c.ShouldBindJSON(&body)
	if err != nil {
		return err
	}
	insert.ItemID = body.ItemID

	selectField := make([]string, 0)

	if body.AgeLimit != nil {
		selectField = append(selectField, "age_limit")
		insert.AgeLimit = *body.AgeLimit
	}

	if body.Location != nil {
		selectField = append(selectField, "location")
		insert.Location = *body.Location
	}

	if body.State != nil {
		selectField = append(selectField, "state")
		insert.State = *body.State
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
	).Where("item_tables.item_id=?", body.ItemID).Updates(&insert).Error
}
