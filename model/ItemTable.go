package model

import (
	"database/sql/driver"
	"errors"
)

type ItemState struct {
	Correct bool `json:"correct"`
	Fixing  bool `json:"fixing"`
	Unlabel bool `json:"unlabel"`
	Discard bool `json:"discard"`
}

//go:generate go run ../tools/generate_router -type "ItemTable" -method "GET" -ignore "ID" -options "State,ItemID"
//go:generate go run ../tools/generate_router -type "ItemTable" -method "POST" -ignore "ID"
//go:generate go run ../tools/generate_router -type "ItemTable" -method "PUT" -ignore "ID,ItemID,Name,Date,Cost" -indexField "ItemID"
type ItemTable struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	ItemID   string    `gorm:"not null;uniqueIndex" json:"item_id"`
	Name     string    `gorm:"not null" json:"name"`
	Date     string    `gorm:"not null" json:"date"`
	AgeLimit uint      `gorm:"not null" json:"age_limit"`
	Cost     uint      `gorm:"not null" json:"cost"`
	Location string    `gorm:"not null" json:"location"`
	State    ItemState `gorm:"not null;type:INT;index" json:"state"`
	Note     string    `gorm:"not null" json:"note"`
}

func (state *ItemState) Scan(value interface{}) error {
	switch value := value.(type) {
	case uint:
		state.Correct = value>>1&1 == 1
		state.Fixing = value>>2&1 == 1
		state.Unlabel = value>>4&1 == 1
		state.Discard = value>>8&1 == 1
	case int:
		state.Correct = value>>1&1 == 1
		state.Fixing = value>>2&1 == 1
		state.Unlabel = value>>4&1 == 1
		state.Discard = value>>8&1 == 1
	case uint64:
		state.Correct = value>>1&1 == 1
		state.Fixing = value>>2&1 == 1
		state.Unlabel = value>>4&1 == 1
		state.Discard = value>>8&1 == 1
	case int64:
		state.Correct = value>>1&1 == 1
		state.Fixing = value>>2&1 == 1
		state.Unlabel = value>>4&1 == 1
		state.Discard = value>>8&1 == 1
	default:
		return errors.New("not a valid type")
	}

	return nil
}

func (state ItemState) Value() (driver.Value, error) {
	var result int64 = 0
	if state.Correct {
		result |= 1 << 1
	}
	if state.Fixing {
		result |= 1 << 2
	}
	if state.Unlabel {
		result |= 1 << 4
	}
	if state.Discard {
		result |= 1 << 8
	}

	return result, nil
}
