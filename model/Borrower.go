package model

//go:generate generate_router -type "Borrower" -method "Update" -ignore "BorrowRecords"
//go:generate generate_router -type "Borrower" -method "Create" -ignore "BorrowRecords"
//go:generate generate_router -type "Borrower" -method "First"
//go:generate generate_router -type "Borrower" -method "Find" -ignore "ID"
type Borrower struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Name          string         `json:"name" gorm:"not null;index"`
	Phone         string         `json:"phone" gorm:"not null;index"`
	BorrowRecords []BorrowRecord `gorm:"foreignKey:ID" json:"borrow_records"`
}
