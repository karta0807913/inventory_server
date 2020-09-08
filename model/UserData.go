package model

//go:generate go run ../tools/generate_router -type "UserData" -method "PUT" -ignore "ID,Account"
type UserData struct {
	ID       uint   `gorm:"primaryKey" json:"user_id"`
	Account  string `gorm:"not null;type=VARCHAR(50);uniqueIndex" json:"account"`
	Password string `gorm:"not null;type=VARCHAR(256)" json:"password"`
	Name     string `gorm:"not null" json:"nickname"`
}
