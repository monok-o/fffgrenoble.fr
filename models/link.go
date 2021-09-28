package models

type Link struct {
	Id    uint   `gorm:"not null;primary_key;autoIncrement" json:"id"`
	Title string `gorm:"not null;" json:"title"`
	Link  string `gorm:"not null;" json:"link"`
}
