package model

//go:generate go run ../tools/generate_router -type "Borrower" -method "Update"
//go:generate go run ../tools/generate_router -type "Borrower" -method "Create"
//go:generate go run ../tools/generate_router -type "Borrower" -method "First"
//go:generate go run ../tools/generate_router -type "Borrower" -method "Find" -ignore "ID"
type Borrower struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name" gorm:"not null;index"`
	Phone string `json:"phone" gorm:"not null;index"`
}
