package service

// TODO: Implement PanicHandler

import (
	"github.com/HermanPlay/web-app-backend/internal/api/http/util"
	"github.com/HermanPlay/web-app-backend/internal/api/http/util/token"
	"github.com/HermanPlay/web-app-backend/internal/config"
	"github.com/HermanPlay/web-app-backend/package/domain/models"
	"github.com/HermanPlay/web-app-backend/package/domain/schemas"
	"github.com/HermanPlay/web-app-backend/package/repository"
	"gorm.io/gorm"
)

const randomPasswordLength = 8

type UserService interface {
	GetAllUser() ([]schemas.User, error)
	GetUserById(userId int) (*schemas.User, error)
	AddUserData(user schemas.UserInput) (*schemas.User, error)
	UpdateUserData(user schemas.UserUpdate, userId int) (*schemas.User, error)
	DeleteUser(userId int) error
	DecodeToken(token string) (*schemas.User, error)
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
	cfg            *config.Config
}

func (u UserServiceImpl) UpdateUserData(user schemas.UserUpdate, userId int) (*schemas.User, error) {
	data, err := u.userRepository.FindUserById(userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	if user.Email != data.Email {
		exists, err := u.userRepository.CheckUserExist(user.Email)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, ErrAlreadyExists
		}
	}

	updateModel(&data, &user)

	updated, err := u.userRepository.Update(&data)
	if err != nil {
		return nil, err
	}
	returnData := createUserSchema(&updated)
	return returnData, nil
}

func (u UserServiceImpl) GetUserById(userId int) (*schemas.User, error) {
	data, err := u.userRepository.FindUserById(userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	returnData := createUserSchema(&data)
	return returnData, nil

}

func (u UserServiceImpl) AddUserData(user schemas.UserInput) (*schemas.User, error) {
	userData := createUserModelWithPassword(&user)
	if userData.Email == "" || userData.Name == "" {
		return nil, ErrInvalidInput
	}

	data, err := u.userRepository.Save(&userData)
	if err != nil {
		return nil, err
	}

	returnData := createUserSchema(&data)

	return returnData, nil
}

func (u UserServiceImpl) GetAllUser() ([]schemas.User, error) {
	data, err := u.userRepository.FindAllUser()
	if err != nil {
		return nil, err
	}
	returnData := make([]schemas.User, 0)
	for _, user := range data {
		returnData = append(returnData, *createUserSchema(&user))
	}
	return returnData, nil

}

func (u UserServiceImpl) DeleteUser(userId int) error {
	_, err := u.userRepository.FindUserById(userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound
		}
		return err
	}
	err = u.userRepository.DeleteUserById(userId)
	if err != nil {
		return err
	}
	return nil
}

func (u UserServiceImpl) DecodeToken(t string) (*schemas.User, error) {
	claims, err := token.DecodeToken(t, u.cfg)
	if err != nil {
		return nil, err
	}

	userId := int(claims["user_id"].(float64))
	if userId == 0 {
		return nil, ErrInvalidToken
	}

	user, err := u.userRepository.FindUserById(userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	returnData := createUserSchema(&user)
	return returnData, nil
}

func createUserSchema(model *models.User) *schemas.User {
	return &schemas.User{
		ID:    model.ID,
		Email: model.Email,
		Name:  model.Name,
		Role:  model.Role,
	}
}

func updateModel(model *models.User, request *schemas.UserUpdate) {
	if request.Name != "" {
		model.Name = request.Name
	}
	if request.Email != "" {
		model.Email = request.Email
	}
	if request.Role != "" {
		model.Role = request.Role
	}
}

func createUserModelWithPassword(requestData *schemas.UserInput) models.User {
	return models.User{
		Name:  requestData.Name,
		Email: requestData.Email,
		Password: util.GenerateRandomString(
			randomPasswordLength,
		),
		Role: models.UserRole, // Default role
	}
}

func NewUserService(userRepository repository.UserRepository, cfg *config.Config) UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
		cfg:            cfg,
	}
}
