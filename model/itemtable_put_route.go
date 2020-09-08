package model

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// this file generate by go generate, please don't edit it
// data will put into struct
func (insert *ItemTable) PUT(c *gin.Context, db *gorm.DB) error {
	type Body struct {
		AgeLimit uint `json:"age_limit" binding:"rqeuired"`

		Location *string    `json:"location"`
		State    *ItemState `json:"state"`
		Note     *string    `json:"note"`
	}
	var body Body
	err := c.ShouldBindJSON(&body)
	if err != nil {
		return err
	}

	selectField := make([]string, 0)

	if body.Location == nil {
		body.Location = new(string)
	} else {
		selectField = append(selectField, "Location")
	}

	if body.State == nil {
		body.State = new(ItemState)
	} else {
		selectField = append(selectField, "State")
	}

	if body.Note == nil {
		body.Note = new(string)
	} else {
		selectField = append(selectField, "Note")
	}

	if len(selectField) == 0 {
		return errors.New("rqeuire at least one option")
	}

	insert.AgeLimit = body.AgeLimit
	insert.Location = *body.Location
	insert.State = *body.State
	insert.Note = *body.Note

	return db.Select(
		selectField[0], selectField[1:],
	).Where("AgeLimit=?", body.AgeLimit).Updates(&insert).Error
}