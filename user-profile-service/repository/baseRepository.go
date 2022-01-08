package repository

import uuid "github.com/satori/go.uuid"

type BaseRepository interface {
	getById(id uuid.UUID) (interface{}, error)
	list(searchParam string, offset int, limit int) ([]interface{}, int, error)
	create(model interface{}) error
	update(model interface{}) error
	delete(model interface{}) error
}
