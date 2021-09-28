package models

type Event struct {
	Id    uint   `gorm:"not null;primary_key;autoIncrement" json:"id"`
	Title string `gorm:"not null;" json:"title"`
	Date  string `gorm:"not null;" json:"date"`
	Time  string `gorm:"not null;" json:"time"`
	Place string `gorm:"not null;" json:"place"`
	Link  string `gorm:"not null;" json:"link"`
}
