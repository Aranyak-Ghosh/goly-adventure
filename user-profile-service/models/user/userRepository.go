package user

import (
	uuid "github.com/satori/go.uuid"
)

func (mw *User) GetById(id uuid.UUID) (*User, error) {
	return nil, nil
}

func (mw *User) List(searchParam string, offset int, limit int) ([]User, int, error)
{
	return nil, 0, nil
}

func (mw *User) Create(model interface{}) error{
	return nil
}

func (mw *User) Update(model interface{}) error{
	return nil
}
func (mw *User) Delete(model interface{}) error{
	return nil
}