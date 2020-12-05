package model

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// this file generate by go generate, please don't edit it
// search options will put into struct
func (item *BorrowRecord) Find(c *gin.Context, db *gorm.DB) ([]BorrowRecord, error) {
	type Body struct {
		ItemID   *uint `form:"item_id"`
		Returned *bool `form:"returned"`
	}
	var body Body
	err := c.ShouldBindQuery(&body)

	whereField := make([]string, 0)
	valueField := make([]interface{}, 0)

	if body.ItemID != nil {
		whereField = append(whereField, "borrow_records.item_id=?")
		valueField = append(valueField, body.ItemID)
		item.ItemID = *body.ItemID
	}

	if body.Returned != nil {
		whereField = append(whereField, "borrow_records.returned=?")
		valueField = append(valueField, body.Returned)
		item.Returned = *body.Returned
	}

	var limit int = 20
	slimit, ok := c.GetQuery("limit")
	if ok {
		limit, err = strconv.Atoi(slimit)
		if err != nil {
			limit = 20
		} else {
			if limit <= 0 || 20 < limit {
				limit = 20
			}
		}
	}
	soffset, ok := c.GetQuery("offset")
	var offset int
	if ok {
		offset, err = strconv.Atoi(soffset)
		if err != nil {
			offset = 0
		} else if offset < 0 {
			offset = 0
		}
	} else {
		offset = 0
	}
	var result []BorrowRecord
	if len(whereField) != 0 {
		db = db.Where(
			strings.Join(whereField, " and "),
			valueField[0], valueField[1:],
		)
	}
	err = db.Limit(limit).Offset(offset).Find(&result).Error
	return result, err
}
