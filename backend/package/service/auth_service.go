package service

import (
	"github.com/HermanPlay/web-app-backend/internal/api/http/util"
	"github.com/HermanPlay/web-app-backend/package/domain/models"
	"github.com/HermanPlay/web-app-backend/package/domain/schemas"
	"github.com/HermanPlay/web-app-backend/package/repository"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	RegisterUser(data schemas.UserRegister) (*schemas.User, error)
	LoginUser(data schemas.UserLogin) (string, error)
	ResetPassword(data schemas.UserResetPassword) (string, error)
}

type AuthServiceImpl struct {
	authRepository repository.AuthRepository
	userRepository repository.UserRepository
}

func (a AuthServiceImpl) RegisterUser(data schemas.UserRegister) (*schemas.User, error) {
	exists, err := a.userRepository.CheckUserExist(data.Email)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	if exists {
		return nil, ErrAlreadyExists
	}

	model := a.createUserModel(&data)
	user, err := a.userRepository.Save(&model)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	returnData := createUserSchema(&user)
	return returnData, nil
}

func (a AuthServiceImpl) LoginUser(data schemas.UserLogin) (string, error) {
	// Skip the check if user exists because it is already checked in the repository
	token, err := a.authRepository.LoginUser(data.Email, data.Password)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", ErrNotFound
		}
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return "", ErrInvalidPassword
		}
		logrus.Error(err)
		return "", err
	}
	return token, nil
}

func (a AuthServiceImpl) ResetPassword(data schemas.UserResetPassword) (string, error) {
	user, err := a.userRepository.GetUserByEmail(data.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", ErrNotFound
		}
		logrus.Error(err)
		return "", err
	}

	new_password := util.GenerateRandomString(8)
	user.Password = new_password
	_, err = a.userRepository.Update(&user)
	if err != nil {
		logrus.Error(err)
		return "", err
	}

	return new_password, nil
}

func (a AuthServiceImpl) createUserModel(requestData *schemas.UserRegister) models.User {
	return models.User{
		Name:     requestData.Name,
		Email:    requestData.Email,
		Password: requestData.Password,
		Role:     models.UserRole, // Default role
	}
}

func NewAuthService(authRepository repository.AuthRepository, userRepository repository.UserRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		authRepository: authRepository,
		userRepository: userRepository,
	}
}
