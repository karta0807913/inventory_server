package model

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (item *ItemTable) Find(c *gin.Context, db *gorm.DB) ([]ItemTable, error) {
	type Body struct {
		Name  *string    `form:"name"`
		State *ItemState `form:"state"`
	}
	var body Body
	err := c.ShouldBindQuery(&body)

	whereField := make([]string, 0)
	valueField := make([]interface{}, 0)

	if body.Name != nil {
		*body.Name = "%" + *body.Name + "%"
		whereField = append(whereField, "item_tables.name like?")
		valueField = append(valueField, body.Name)
		item.Name = *body.Name
	}

	if body.State != nil {
		whereField = append(whereField, "item_tables.state=?")
		valueField = append(valueField, body.State)
		item.State = *body.State
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
	var result []ItemTable
	if len(whereField) != 0 {
		db = db.Where(
			strings.Join(whereField, " and "),
			valueField[0], valueField[1:],
		)
	}
	err = db.Limit(limit).Offset(offset).Find(&result).Error
	return result, err
}
