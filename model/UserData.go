package model

//go:generate generate_router -type "UserData" -method "Update" -ignore "ID,Account"
type UserData struct {
	//使用者ID
	ID uint `gorm:"primaryKey" json:"user_id"`
	//使用者帳號
	Account string `gorm:"not null;type=VARCHAR(50);uniqueIndex" json:"account"`
	//使用者密碼
	Password string `gorm:"not null;type=VARCHAR(256)" json:"password"`
	//使用者名稱
	Name string `gorm:"not null" json:"nickname"`
}
