package model

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// this file generate by go generate, please don't edit it
// data will put into struct
func (insert *ItemTable) Create(c *gin.Context, db *gorm.DB) error {
	type Body struct {
		ItemID   string    `json:"item_id" binding:"required"`
		Name     string    `json:"name" binding:"required"`
		Date     string    `json:"date" binding:"required"`
		AgeLimit uint      `json:"age_limit" binding:"required"`
		Cost     uint      `json:"cost" binding:"required"`
		Location string    `json:"location" binding:"required"`
		State    ItemState `json:"state" binding:"required"`
		Note     string    `json:"note" binding:"required"`
	}
	var body Body
	err := c.ShouldBindJSON(&body)
	if err != nil {
		return err
	}

	selectField := []string{
		"item_tables.item_id",
		"item_tables.name",
		"item_tables.date",
		"item_tables.age_limit",
		"item_tables.cost",
		"item_tables.location",
		"item_tables.state",
		"item_tables.note",
	}

	insert.ItemID = body.ItemID
	insert.Name = body.Name
	insert.Date = body.Date
	insert.AgeLimit = body.AgeLimit
	insert.Cost = body.Cost
	insert.Location = body.Location
	insert.State = body.State
	insert.Note = body.Note

	return db.Select(
		selectField[0], selectField[1:],
	).Create(&insert).Error
}
