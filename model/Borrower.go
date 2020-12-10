package model

//go:generate generate_router -type "Borrower" -method "Update" -ignore "BorrowRecords"
//go:generate generate_router -type "Borrower" -method "Create" -ignore "BorrowRecords"
//go:generate generate_router -type "Borrower" -method "First"
//go:generate generate_router -type "Borrower" -method "Find" -ignore "ID"
type Borrower struct {
	//借貸人ID
	ID uint `gorm:"primaryKey" json:"id"`
	//借貸人名稱
	Name string `json:"name" gorm:"not null;index"`
	//借貸人手機
	Phone string `json:"phone" gorm:"not null;index"`
	//借貸紀錄
	BorrowRecords []BorrowRecord `gorm:"foreignKey:BorrowerID" json:"borrow_records"`
}
