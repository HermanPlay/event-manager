package repository

import (
	"testing"

	"github.com/HermanPlay/web-app-backend/internal/config"
	"github.com/HermanPlay/web-app-backend/package/domain/models"
	"github.com/HermanPlay/web-app-backend/package/utils"
	"gorm.io/gorm"
)

func TestNewAuthRepository(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	cfg := config.Config{
		Db:  config.Db{},
		App: config.App{},
	}
	authRepositoryImpl := NewAuthRepository(db, &cfg)
	if authRepositoryImpl == nil {
		t.Errorf("Auth repository is nil, when not expected")
	}
}

func TestLoginUser(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	cfg := config.Config{
		Db:  config.Db{},
		App: config.App{},
	}
	authRepositoryImpl := NewAuthRepository(db, &cfg)
	password := "passwordlong"
	want := models.User{
		Email:    "email@email.com",
		Password: password,
		Role:     models.UserRole,
	}
	userRepository, err := NewUserRepository(db)
	if err != nil {
		t.Errorf("Error when create new user repository, when not expected. Error: %v", err)
	}
	_, err = userRepository.Save(&want)
	if err != nil {
		t.Errorf("Error when save user, when not expected. Error: %v", err)
	}
	t.Run("Invalid email", func(t *testing.T) {
		got, err := authRepositoryImpl.LoginUser("invalid", password)
		if err != gorm.ErrRecordNotFound {
			t.Errorf("Error is not gorm.ErrRecordNotFound, when expected. Error: %v", err)
		}
		if got != "" {
			t.Errorf("Token is not empty, when expected, got: %v", got)
		}

	})
	t.Run("Invalid password", func(t *testing.T) {
		_, err := authRepositoryImpl.LoginUser(want.Email, "invalid")
		if err == nil {
			t.Errorf("Error is nil, when expected")
		}
	})
	t.Run("Valid login", func(t *testing.T) {
		token, err := authRepositoryImpl.LoginUser(want.Email, password)
		if err != nil {
			t.Errorf("Error when login, when not expected. Error: %v", err)
			return
		}
		if token == "" {
			t.Errorf("Token is empty, when not expected")
		}
	})

}
