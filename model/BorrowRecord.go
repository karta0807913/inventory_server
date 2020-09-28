package model

import "time"

//go:generate go run ../tools/generate_router -type "BorrowRecord" -method "Find" -ignore "Borrower,ID,BorrowDate"
//go:generate go run ../tools/generate_router -type "BorrowRecord" -method "First" -ignore "Borrower,BorrowDate"
//go:generate go run ../tools/generate_router -type "BorrowRecord" -method "Create" -options "Borrower"
//go:generate go run ../tools/generate_router -type "BorrowRecord" -method "Update" -ignore "Borrower"
type BorrowRecord struct {
	Borrower   Borrower  `json:"borrower" gorm:"foreignKey:borrower_id"`
	ID         uint      `gorm:"primaryKey" json:"id"`
	BorrowerID uint      `gorm:"not null" json:"borrower_id"`
	BorrowDate time.Time `gorm:"index;not null" json:"borrow_date"`
	ReplyDate  time.Time `gorm:"index" json:"reply_date"`
	Note       string    `json:"note"`
}
