package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	name        string
	email       string
	imgUrl      string
	country     string
	dateOfBirth time.Time

	follows []User `gorm:"many2many:user_follows"`
}
