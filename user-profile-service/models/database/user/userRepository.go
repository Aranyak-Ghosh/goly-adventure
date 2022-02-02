package user

import (
	"time"

	"github.com/Aranyak-Ghosh/spotigo/user_profile/errors"
	"go.uber.org/fx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (mw *UserRepository) GetById(id string) (User, error) {

	err := validateId(id)

	if err != nil {
		return User{}, err
	}

	var user User

	result := mw.db.Where("id = ?", id).First(&user)

	return user, result.Error
}

func (mw *UserRepository) List(searchParam string, offset int, limit int) ([]User, int64, error) {
	var users []User
	var count int64
	query := mw.db.Model(&users).Where("name LIKE ?", "%"+searchParam+"%").Or("email LIKE ?", "%"+searchParam+"%")
	query.Count(&count)
	queryResult := query.Order("name DESC").Offset(offset).Limit(limit).Find(&users)
	return users, count, queryResult.Error
}

func (mw *UserRepository) Create(model *User) error {
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

func (mw *UserRepository) Update(model *User) error {
	result := mw.db.Save(&model)
	return result.Error
}

func (mw *UserRepository) Delete(model *User) error {
	result := mw.db.Delete(&model)
	return result.Error
}

func (mw *UserRepository) Follow(user *User, follow *User) error {
	err := mw.db.Model(&user).Association("Follows").Append(follow)
	return err
}

func (mw *UserRepository) UnFollow(user *User, follow *User) error {
	err := mw.db.Model(&user).Association("Follows").Delete(follow)
	return err
}

func (mw *UserRepository) ListFollowing(userId string, search string, offset int, limit int) ([]User, int64, error) {
	var users []User
	var count int64

	query := mw.db.Model(&User{ID: userId}).Where("name LIKE ?", "%"+search+"%").Or("email LIKE ?", "%"+search+"%").Association("Follows")

	count = query.Count()
	queryError := query.Find(&users)

	if int64(offset) > count {
		offset = int(count)
	}

	end := offset + limit
	if int64(end) > count {
		end = int(count)
	}
	return users[offset:end], count, queryError
}

func (mw *UserRepository) ListFollowers(userId string, search string, offset int, limit int) ([]User, int64, error) {

	var followers []User
	var count int64
	query := mw.db.Joins("inner JOIN user_follows on user_follows.user_id = users.id").Where("follow_id = ?", userId).Where("users.name LIKE ? or users.email LIKE ?", "%"+search+"%", "%"+search+"%").Model(&User{})

	_ = query.Count(&count)
	queryError := query.Limit(limit).Offset(offset).Find(&followers)

	return followers, count, queryError.Error
}

func validateId(id string) error {
	if len(id) != 36 {
		return errors.ErrUUIDValidationFailed
	}
	var err error
	_, err = uuid.Parse(id)
	if err != nil {
		return errors.ErrUUIDValidationFailed
	}

	return nil
}

func (mw *UserRepository) Migrate() error {
	result := mw.db.AutoMigrate(&User{})
	return result
}

func (user *UserRepository) SeedData() error {
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

func NewUserRepository(db *gorm.DB) (user *UserRepository) {
	user = &UserRepository{db: db}
	user.Migrate()
	return user
}

var Module = fx.Option(fx.Provide(NewUserRepository))
