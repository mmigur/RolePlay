package models

import "gorm.io/gorm"

type CodeForEmail struct {
	gorm.Model
	Email string
	Code  string
}
