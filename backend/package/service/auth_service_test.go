package service

import (
	"testing"

	"github.com/HermanPlay/web-app-backend/internal/config"
	"github.com/HermanPlay/web-app-backend/package/domain/models"
	"github.com/HermanPlay/web-app-backend/package/domain/schemas"
	"github.com/HermanPlay/web-app-backend/package/repository"
	"github.com/HermanPlay/web-app-backend/package/utils"
)

func TestRegisterUser(t *testing.T) {

	db := utils.ConnectToTestDatabase()

	cfg := config.Config{
		Db:  config.Db{},
		App: config.App{ApiSecret: "secret"},
	}
	authRepository := repository.NewAuthRepository(db, &cfg)
	userRepository, err := repository.NewUserRepository(db)
	if err != nil {
		t.Errorf("Error when create new user repository, when not expected. Error: %v", err)
	}

	authService := NewAuthService(authRepository, userRepository)
	t.Run("correct user", func(t *testing.T) {
		user := schemas.UserRegister{
			Name:     "name",
			Email:    "email",
			Password: "password",
		}
		registered, err := authService.RegisterUser(user)
		if err != nil {
			t.Errorf("Error when register user, when not expected. Error: %v", err)
		}
		compareUser(t, *registered, models.User{Name: user.Name, Email: user.Email, Role: models.UserRole})
	})
	t.Run("wrong user", func(t *testing.T) {
		// Create existing user
		user := &models.User{
			Name:     "name",
			Email:    "email",
			Password: "password",
		}
		userRepository.Save(user)

		// Try to register user with the same email
		userRegister := schemas.UserRegister{
			Name:     "name",
			Email:    "email",
			Password: "password",
		}
		registered, err := authService.RegisterUser(userRegister)
		if err == nil {
			t.Errorf("Error is nil, when expected")
		}
		if err != ErrAlreadyExists {
			t.Errorf("Error is not ErrAlreadyExists, when expected. Error: %v", err)
		}
		if registered != nil {
			t.Errorf("Registered is not nil, when expected")
		}

	})
}

func TestLoginUser(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	cfg := config.Config{
		Db:  config.Db{},
		App: config.App{},
	}
	authRepository := repository.NewAuthRepository(db, &cfg)
	userRepository, err := repository.NewUserRepository(db)
	if err != nil {
		t.Errorf("Error when create new user repository, when not expected. Error: %v", err)
	}

	authService := NewAuthService(authRepository, userRepository)
	password := "passwordlong"
	want := models.User{
		Email:    "email",
		Password: password,
		Role:     models.UserRole,
	}
	userRepository.Save(&want)
	t.Run("Invalid email", func(t *testing.T) {
		userLogin := schemas.UserLogin{
			Email:    "invalid",
			Password: password,
		}
		got, err := authService.LoginUser(userLogin)
		if err != ErrNotFound {
			t.Errorf("Error is not ErrNotFound, when expected. Error: %v", err)
		}
		if got != "" {
			t.Errorf("Token is not empty, when expected, got: %v", got)
		}
	})
	t.Run("Invalid password", func(t *testing.T) {
		userLogin := schemas.UserLogin{
			Email:    want.Email,
			Password: "invalid",
		}
		got, err := authService.LoginUser(userLogin)
		if err != ErrInvalidPassword {
			t.Errorf("Error is not ErrInvalidPassword, when expected")
		}
		if got != "" {
			t.Errorf("Token is not empty, when expected, got: %v", got)
		}
	})
	t.Run("Valid login", func(t *testing.T) {
		userLogin := schemas.UserLogin{
			Email:    want.Email,
			Password: password,
		}
		token, err := authService.LoginUser(userLogin)
		if err != nil {
			t.Errorf("Error when login, when not expected. Error: %v", err)
			return
		}
		if token == "" {
			t.Errorf("Token is empty, when not expected")
		}
	})
}

func TestResetPassword(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	cfg := config.Config{
		Db:  config.Db{},
		App: config.App{},
	}
	authRepository := repository.NewAuthRepository(db, &cfg)
	userRepository, err := repository.NewUserRepository(db)
	if err != nil {
		t.Errorf("Error when create new user repository, when not expected. Error: %v", err)
	}

	authService := NewAuthService(authRepository, userRepository)

	// Create user
	user := models.User{
		Email:    "email",
		Password: "password",
		Role:     models.UserRole,
	}
	userRepository.Save(&user)

	t.Run("Invalid email", func(t *testing.T) {
		resetPassword := schemas.UserResetPassword{
			Email: "invalid",
		}
		_, err := authService.ResetPassword(resetPassword)
		if err != ErrNotFound {
			t.Errorf("Error is not ErrNotFound, when expected. Error: %v", err)
		}
	})
	t.Run("Valid email", func(t *testing.T) {
		resetPassword := schemas.UserResetPassword{
			Email: user.Email,
		}
		_, err := authService.ResetPassword(resetPassword)
		if err != nil {
			t.Errorf("Error when reset password, when not expected. Error: %v", err)
		}
	})
}
