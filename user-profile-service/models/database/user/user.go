package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          string `gorm:"primary_key;size:36"`
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

type UserRepository struct {
	db *gorm.DB
}
