package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
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

func (mw *UserDAO) GetById(id uuid.UUID) (User, error) {
	var user User
	result := mw.db.First(user, id.String())

	return user, result.Error
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
	if model.ID == "" {
		model.ID = uuid.NewString()
	}
	err := validateId(model.ID)
	if err != nil {
		return err
	}
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

func (mw *UserDAO) ListFollowing(userId string, search string, offset int, limit int) ([]User, int64, error) {
	var user = User{ID: userId}
	var users []User
	var count *int64

	query := mw.db.Joins("JOIN user_follows on user_follows.user_id = users.id").Joins("JOIN users as following on user_follows.follow_id = following.id").Preload("Follows").Where("following.name LIKE ?", "%"+search+"%").Model(&user)

	queryError := query.Limit(limit).Offset(offset).Find(&users)
	_ = query.Count(count)
	return users, *count, queryError.Error
}

func (mw *UserDAO) ListFollowers(userId string, search string, offset int, limit int) ([]User, int64, error) {

	var followers []User
	var count int64
	query := mw.db.Joins("JOIN user_follows on user_follows.user_id = users.id").Joins("JOIN users as following on user_follows.follow_id = following.id").Preload("Follows").Where("following.ID = ?", userId).Where("users.id LIKE ?", "%"+search+"%").Model(&User{})

	_ = query.Count(&count)
	queryError := query.Limit(limit).Offset(offset).Find(followers)

	return followers, count, queryError.Error
}

func validateId(id string) error {
	if len(id) != 36 {
		return errors.New("invalid id")
	}
	var err error
	_, err = uuid.Parse(id)
	if err != nil {
		return err
	}

	return nil
}
