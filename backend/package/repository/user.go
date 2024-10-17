package repository

import (
	"github.com/HermanPlay/web-app-backend/package/domain/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAllUser() ([]models.User, error)
	FindUserById(id int) (models.User, error)
	Save(user *models.User) (models.User, error)
	DeleteUserById(id int) error
	CheckUserExist(email string) (bool, error)
	Update(user *models.User) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (u UserRepositoryImpl) FindAllUser() ([]models.User, error) {
	var users []models.User

	var err = u.db.Find(&users).Error
	if err != nil {
		logrus.Error("Got an error finding all couples. Error: ", err)
		return nil, err
	}

	return users, nil
}

func (u UserRepositoryImpl) FindUserById(id int) (models.User, error) {
	user := models.User{
		ID: id,
	}
	err := u.db.First(&user).Error
	if err != nil {
		logrus.Error("Got and error when find user by id. Error: ", err)
		return models.User{}, err
	}
	return user, nil
}

func (u UserRepositoryImpl) Save(user *models.User) (models.User, error) {
	err := u.db.Save(user).Error
	if err != nil {
		logrus.Error("Got an error when save user. Error: ", err)
		return models.User{}, err
	}
	return *user, nil
}

func (u UserRepositoryImpl) DeleteUserById(id int) error {
	err := u.db.Delete(&models.User{}, id).Error
	if err != nil {
		logrus.Error("Got an error when delete user. Error: ", err)
		return err
	}
	return nil
}

func (u UserRepositoryImpl) CheckUserExist(email string) (bool, error) {
	var user models.User
	err := u.db.Model(&user).Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (u UserRepositoryImpl) Update(user *models.User) (models.User, error) {
	err := u.db.Save(user).Error
	if err != nil {
		logrus.Error("Got an error when update user. Error: ", err)
		return models.User{}, err
	}
	return *user, nil
}

func (u UserRepositoryImpl) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := u.db.Model(&user).Where("email = ?", email).First(&user).Error
	if err != nil {
		logrus.Error("Got an error when get user by email. Error: ", err)
		return models.User{}, err
	}
	return user, nil
}

func NewUserRepository(db *gorm.DB) (*UserRepositoryImpl, error) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}
	return &UserRepositoryImpl{
		db: db,
	}, nil
}
