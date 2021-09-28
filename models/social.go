package models

type Social struct {
	Id   uint   `gorm:"not null;primary_key;autoIncrement" json:"id"`
	Icon string `gorm:"not null;" json:"icon"`
	Link string `gorm:"not null;" json:"link"`
}
