package models

type Category struct {
	Id   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
}
