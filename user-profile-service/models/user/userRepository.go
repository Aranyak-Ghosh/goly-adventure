package user

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func (user *UserDAO) Initialize(db *gorm.DB) {
	user.db = db
}
func (user *UserDAO) SeedData() error {
	err := user.Create(&User{
		Name:        "Aranyak Ghosh",
		Email:       "aranyakghosh@gmail.com",
		Country:     "India",
		Public:      true,
		DateOfBirth: time.Date(1996, time.December, 16, 0, 0, 0, 0, time.UTC),
	})

	user.Create(&User{
		Name:        "Shailika Garg",
		Email:       "shailika.garg@gmail.com",
		Country:     "India",
		Public:      true,
		DateOfBirth: time.Date(1992, time.August, 21, 0, 0, 0, 0, time.UTC),
	})

	return err
}

func (mw *UserDAO) GetById(id uuid.UUID) (*User, error) {
	result := mw.db.First(&mw.User, id)

	return &mw.User, result.Error
}

func (mw *UserDAO) List(searchParam string, offset int, limit int) ([]User, int64, error) {
	var users []User
	var count int64
	query := mw.db.Model(&users).Where("name LIKE ?", "%"+searchParam+"%").Or("email LIKE ?", "%"+searchParam+"%")
	query.Count(&count)
	queryResult := query.Order("name DESC").Offset(offset).Limit(limit).Find(&users)
	return users, count, queryResult.Error
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
