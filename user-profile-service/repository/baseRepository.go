package repository

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type BaseRepository interface {
	Initialize(db *gorm.DB)
	GetById(id uuid.UUID) (*interface{}, error)
	List(searchParam string, offset int, limit int) ([]interface{}, int, error)
	Create(model interface{}) error
	Update(model interface{}) error
	Delete(model interface{}) error
	Migrate() error
}
