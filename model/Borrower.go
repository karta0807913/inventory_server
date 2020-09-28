package model

//go:generate go run ../tools/generate_router -type "Borrower" -method "Update"
type Borrower struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name" gorm:"not null"`
	Phone string `json:"phone" gorm:"not null;uniqueIndex"`
}
