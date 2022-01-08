package user

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func (user *UserDAO) Initialize(db *gorm.DB) {
	user.db = db
}

func (mw *UserDAO) GetById(id uuid.UUID) (*User, error) {
	result := mw.db.First(&mw.User, id)

	return &mw.User, result.Error
}

func (mw *UserDAO) List(searchParam string, offset int, limit int) ([]User, int64, error) {
	var users []User
	var count *int64
	query := mw.db.Where("name LIKE ?", "%"+searchParam+"%").Or("email LIKE ?", "%"+searchParam+"%").Order("name DESC")
	queryResult := query.Offset(offset).Limit(limit).Find(&users)
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

func (mw *UserDAO) Migrate() error {
	result := mw.db.AutoMigrate(&User{})
	return result
}

func (mw *UserDAO) Follow(user *User, follow *User) error {
	err := mw.db.Model(&user).Association("Follows").Append(follow)
	return err
}

func (mw *UserDAO) UnFollow(user *User, follow *User) error {
	err := mw.db.Model(&user).Association("Follows").Delete(follow)
	return err
}

func (mw *UserDAO) ListFollowing(user *User, offset int, limit int) ([]User, int64, error) {
	var users []User
	var count *int64
	query := mw.db.Model(&user).Association("Follows")
	queryError := query.Find(&users)
	*count = query.Count()
	return users, *count, queryError
}

// func (mw *UserDAO) ListFollowers(user *U)
