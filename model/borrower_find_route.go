package model

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// this file generate by go generate, please don't edit it
// search options will put into struct
func (item *Borrower) Find(c *gin.Context, db *gorm.DB) ([]Borrower, error) {
	type Body struct {
		Name  *string `form:"name"`
		Phone *string `form:"phone"`
	}
	var body Body
	err := c.ShouldBindQuery(&body)

	whereField := make([]string, 0)
	valueField := make([]interface{}, 0)

	if body.Name != nil {
		whereField = append(whereField, "borrowers.name=?")
		valueField = append(valueField, body.Name)
		item.Name = *body.Name
	}

	if body.Phone != nil {
		whereField = append(whereField, "borrowers.phone=?")
		valueField = append(valueField, body.Phone)
		item.Phone = *body.Phone
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
	offset, err := strconv.Atoi(soffset)
	if err != nil {
		offset = 0
	} else if offset < 0 {
		offset = 0
	}
	var result []Borrower
	if len(whereField) != 0 {
		db = db.Where(
			strings.Join(whereField, "and"),
			valueField,
		)
	}
	err = db.Limit(limit).Offset(offset).Find(&result).Error
	return result, err
}
