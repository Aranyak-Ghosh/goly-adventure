package user

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	id          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	name        string
	email       string `gorm: "unique_index"`
	imgUrl      string
	country     string
	public      bool `gorm:"default:true"`
	dateOfBirth time.Time
	createdAt   time.Time
	updatedAt   time.Time
	deletedAt   gorm.DeletedAt `gorm:"index"`

	follows []User `gorm:"many2many:user_follows"`
}

type UserDAO struct {
	db *gorm.DB
	User
}
