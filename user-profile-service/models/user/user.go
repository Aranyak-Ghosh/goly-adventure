package user

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uniqueidentifier;primary_key;default:NEWID()"`
	Name        string
	Email       string `gorm: "unique_index"`
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
