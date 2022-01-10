package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          string `gorm:"primary_key;"`
	Name        string
	Email       string `gorm:"unique_index"`
	ImgUrl      string
	Country     string
	Public      bool `gorm:"default:true"`
	DateOfBirth time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	Follows []User `gorm:"many2many:user_follows"`
}

type UserDAO struct {
	db *gorm.DB
	User
}
