package model

import (
	"time"
)

//go:generate generate_router -type "BorrowRecord" -method "Find" -ignore "BorrowDate,ReplyDate,ID"
//go:generate generate_router -type "BorrowRecord" -method "First" -ignore "ItemID,BorrowDate,Returned"
//go:generate generate_router -type "BorrowRecord" -method "Create" -options "Note,ReplyDate" -ignore "Returned,Item"
//go:generate generate_router -type "BorrowRecord" -method "Update" -ignore "Item,ItemID,BorrowDate"
type BorrowRecord struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	ItemID     uint      `gorm:"not null;index" json:"item_id"`
	BorrowDate time.Time `gorm:"not null" json:"borrow_date"`
	ReplyDate  time.Time `json:"reply_date"`
	Note       string    `json:"note"`
	Returned   bool      `json:"returned" gorm:"index;default:false;not null"`
	Item       ItemTable `json:"item" gorm:"foreignKey:ItemID"`
}
