package models

import "time"

type User struct {
	ID         uint `gorm:"primaryKey"`
	Username   string
	Email      string `gorm:"unique;not null"`
	Password   string
	FirstName  string
	MiddleName string
	LastName   string
	Address    string
	CreatedAt  time.Time
}
