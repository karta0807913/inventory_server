package model

type ItemTable struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	ItemID   string `gorm:"not null;uniqueIndex" json:"item_id"`
	Name     string `gorm:"not null" json:"name"`
	Date     string `gorm:"not null" json:"date"`
	AgeLimit uint   `gorm:"not null" json:"age_limit"`
}
