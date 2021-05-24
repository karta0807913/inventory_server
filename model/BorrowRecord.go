package model

import (
	"time"
)

//go:generate generate_router -type "BorrowRecord" -method "Find" -ignore "BorrowDate,ReplyDate,ID"
//go:generate generate_router -type "BorrowRecord" -method "First" -ignore "ItemID,BorrowDate,Returned"
//go:generate generate_router -type "BorrowRecord" -method "Create" -options "Note,ReplyDate" -ignore "Returned,Item" -minItem 0
type BorrowRecord struct {
	//借貸紀錄ID
	ID uint `gorm:"primaryKey" json:"id"`
	//借出人ID
	BorrowerID uint `gorm:"index;not null" json:"borrower_id"`
	//借出物品ID
	ItemID uint `gorm:"not null;index" json:"item_id"`
	//借出時間
	BorrowDate time.Time `gorm:"not null" json:"borrow_date"`
	//收回物品時間
	ReplyDate *time.Time `json:"reply_date"`
	//備註
	Note string `json:"note"`
	//是否歸還
	Returned bool `json:"returned" gorm:"index;default:false;not null"`
	//借出物品詳細
	Item ItemTable `json:"item" gorm:"foreignKey:ItemID"`
}
