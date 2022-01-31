package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          string         `gorm:"primary_key;size:36" json:"id"`
	Name        string         `json:"name"`
	Email       string         `gorm:"unique_index" json:"email"`
	ImgUrl      string         `json:"imgUrl"`
	Country     string         `json:"country"`
	Public      bool           `gorm:"default:true" json:"public"`
	DateOfBirth time.Time      `json:"dateOfBirth"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	Follows []User `gorm:"many2many:user_follows" json:"follows,omitempty"`
}

type UserRepository struct {
	db *gorm.DB
}
