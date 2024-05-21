package models

type CodeForEmail struct {
	Id        uint `gorm:"primaryKey"`
	Email     string
	Code      string
	CreatedAt string
}
