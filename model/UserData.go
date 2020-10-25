package model

//go:generate generate_router -type "UserData" -method "Update" -ignore "ID,Account"
type UserData struct {
	ID       uint   `gorm:"primaryKey" json:"user_id"`
	Account  string `gorm:"not null;type=VARCHAR(50);uniqueIndex" json:"account"`
	Password string `gorm:"not null;type=VARCHAR(256)" json:"password"`
	Name     string `gorm:"not null" json:"nickname"`
}
