package model

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
)

//go:generate go run ../tools/generate_router -type "BorrowRecord" -method "Find" -ignore "Borrower,ID,BorrowDate,ReplyDate"
//go:generate go run ../tools/generate_router -type "BorrowRecord" -method "First" -ignore "Borrower,BorrowDate"
//go:generate go run ../tools/generate_router -type "BorrowRecord" -method "Create" -options "Borrower,BorrowerID,Note" -ignore "Returned"
//go:generate go run ../tools/generate_router -type "BorrowRecord" -method "Update" -ignore "Borrower"
type BorrowRecord struct {
	Borrower   Borrower  `json:"borrower" gorm:"foreignKey:borrower_id"`
	ID         uint      `gorm:"primaryKey" json:"id"`
	BorrowerID uint      `gorm:"not null" json:"borrower_id"`
	BorrowDate time.Time `gorm:"index;not null" json:"borrow_date"`
	ReplyDate  time.Time `gorm:"index" json:"reply_date"`
	Note       string    `json:"note"`
	Returned   bool      `json:"returned" gorm:"index;default:false;not null"`
}

func (record BorrowRecord) MarshalJSON() ([]byte, error) {
	result := gin.H{
		"id":          record.ID,
		"borrower_id": record.BorrowerID,
		"borrow_date": record.BorrowDate,
		"reply_date":  record.ReplyDate,
		"note":        record.Note,
		"returned":    record.Returned,
	}
	if record.Borrower.ID != 0 {
		result["borrower"] = record.Borrower
	}
	return json.Marshal(result)
}
