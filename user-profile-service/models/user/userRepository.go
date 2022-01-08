package user

import (
	uuid "github.com/satori/go.uuid"
)

func (mw *UserDAO) GetById(id uuid.UUID) (*User, error) {
	result := mw.db.First(&mw.User, id)

	return &mw.User, result.Error
}

func (mw *UserDAO) List(searchParam string, offset int, limit int) ([]User, int64, error) {
	var users []User
	var count *int64
	query := mw.db.Where("name LIKE ?", "%"+searchParam+"%").Or("email LIKE ?", "%"+searchParam+"%").Order("name DESC").Offset(offset).Limit(limit)
	queryResult := query.Find(&users)
	query.Count(count)
	return users, *count, queryResult.Error
}

func (mw *UserDAO) Create(model *User) error {
	result := mw.db.Create(&model)
	return result.Error
}

func (mw *UserDAO) Update(model *User) error {
	result := mw.db.Save(&model)
	return result.Error
}
func (mw *UserDAO) Delete(model *User) error {
	result := mw.db.Delete(&model)
	return result.Error
}
