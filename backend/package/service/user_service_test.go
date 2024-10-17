package service

import (
	"testing"

	"github.com/HermanPlay/web-app-backend/internal/api/http/util/token"
	"github.com/HermanPlay/web-app-backend/internal/config"
	"github.com/HermanPlay/web-app-backend/package/domain/models"
	"github.com/HermanPlay/web-app-backend/package/domain/schemas"
	"github.com/HermanPlay/web-app-backend/package/repository"
	"github.com/HermanPlay/web-app-backend/package/utils"
)

func TestGetAll(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	cfg := config.Config{
		Db:  config.Db{},
		App: config.App{ApiSecret: "secret"},
	}
	userRepository, err := repository.NewUserRepository(db)
	if err != nil {
		t.Errorf("Error when create new user repository, when not expected. Error: %v", err)
	}
	userService := NewUserService(userRepository, &cfg)
	t.Run("Empty users", func(t *testing.T) {
		users, err := userService.GetAllUser()
		if err != nil {
			t.Errorf("Error when get all users, when not expected. Error: %v", err)
		}
		if len(users) != 0 {
			t.Errorf("Users is not empty, when expected")
		}
	})
	t.Run("One user", func(t *testing.T) {
		want := models.User{
			Name:     "name",
			Email:    "email",
			Password: "password",
			Role:     "role",
		}
		_, err := userRepository.Save(&want)
		if err != nil {
			t.Errorf("Error when save user, when not expected. Error: %v", err)
		}
		users, err := userService.GetAllUser()
		if err != nil {
			t.Errorf("Error when get all users, when not expected. Error: %v", err)
		}
		if len(users) != 1 {
			t.Errorf("Users is not empty, when expected")
		}
		compareUser(t, users[0], want)

	})
}

func TestGetUserById(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	cfg := config.Config{
		Db:  config.Db{},
		App: config.App{ApiSecret: "secret"},
	}
	userRepository, err := repository.NewUserRepository(db)
	if err != nil {
		t.Errorf("Error when create new user repository, when not expected. Error: %v", err)
	}
	userService := NewUserService(userRepository, &cfg)
	t.Run("Empty user", func(t *testing.T) {
		user, err := userService.GetUserById(1)
		if err == nil {
			t.Errorf("Error is nil, when expected")
		}
		if err != ErrNotFound {
			t.Errorf("Error is not ErrNotFound, when expected")
		}
		if user != nil {
			t.Errorf("User is not nil, when expected")
		}
	})
	t.Run("Existing user", func(t *testing.T) {
		want := models.User{
			Name:     "name",
			Email:    "email",
			Password: "password",
			Role:     "role",
		}
		got, _ := userRepository.Save(&want)
		user, err := userService.GetUserById(got.ID)
		if err != nil {
			t.Errorf("Error when get user by id, when not expected. Error: %v", err)
		}
		compareUser(t, *user, want)
	})
}

func TestAddUserData(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	cfg := config.Config{
		Db:  config.Db{},
		App: config.App{ApiSecret: "secret"},
	}
	userRepository, err := repository.NewUserRepository(db)
	if err != nil {
		t.Errorf("Error when create new user repository, when not expected. Error: %v", err)
	}
	userService := NewUserService(userRepository, &cfg)
	t.Run("Empty user", func(t *testing.T) {
		user, err := userService.AddUserData(schemas.UserInput{})
		if err == nil {
			t.Errorf("Error is nil, when expected")
		}
		if user != nil {
			t.Errorf("User is not nil, when expected")
		}
	})
	t.Run("Valid user", func(t *testing.T) {
		want := schemas.UserInput{
			Name:  "name",
			Email: "email",
			Role:  models.UserRole,
		}
		user, err := userService.AddUserData(want)
		if err != nil {
			t.Errorf("Error when add user data, when not expected. Error: %v", err)
		}
		if user == nil {
			t.Errorf("User is nil, when not expected")
		}
		compareUser(t, *user, models.User{Name: want.Name, Email: want.Email, Role: want.Role})
	})
}

func TestUpdateUser(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	cfg := config.Config{
		Db:  config.Db{},
		App: config.App{ApiSecret: "secret"},
	}
	userRepository, err := repository.NewUserRepository(db)
	if err != nil {
		t.Errorf("Error when create new user repository, when not expected. Error: %v", err)
	}
	userService := NewUserService(userRepository, &cfg)
	t.Run("Non existing user", func(t *testing.T) {
		user, err := userService.UpdateUserData(schemas.UserUpdate{}, 1)
		if err == nil {
			t.Errorf("Error is nil, when expected")
		}
		if err != ErrNotFound {
			t.Errorf("Error is not ErrNotFound, when expected")
		}
		if user != nil {
			t.Errorf("User is not nil, when expected")
		}

	})

	t.Run("Existing user", func(t *testing.T) {
		want := models.User{
			Name:     "name",
			Email:    "email",
			Password: "password",
			Role:     "role",
		}
		want2 := models.User{
			Name:     "name",
			Email:    "existing@email",
			Password: "password",
			Role:     "role",
		}
		got, _ := userRepository.Save(&want)
		userRepository.Save(&want2)
		t.Run("Invalid email", func(t *testing.T) {
			user, err := userService.UpdateUserData(schemas.UserUpdate{Email: want2.Email}, got.ID)
			if err == nil {
				t.Errorf("Error is nil, when expected")
			}
			if err != ErrAlreadyExists {
				t.Errorf("Error is not ErrAlreadyExists, when expected")
			}
			if user != nil {
				t.Errorf("User is not nil, when expected")
			}
		})
		t.Run("Valid email", func(t *testing.T) {
			user, err := userService.UpdateUserData(schemas.UserUpdate{Email: "email2"}, got.ID)
			if err != nil {
				t.Errorf("Error when update user data, when not expected. Error: %v", err)
			}
			if user == nil {
				t.Errorf("User is nil, when not expected")
			}
			compareUser(t, *user, models.User{Name: want.Name, Email: "email2", Role: want.Role})
		})
	})
}

func TestDeleteUser(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	cfg := config.Config{
		Db:  config.Db{},
		App: config.App{ApiSecret: "secret"},
	}
	userRepository, err := repository.NewUserRepository(db)
	if err != nil {
		t.Errorf("Error when create new user repository, when not expected. Error: %v", err)
	}
	userService := NewUserService(userRepository, &cfg)
	t.Run("Non existing user", func(t *testing.T) {
		err := userService.DeleteUser(1)
		if err == nil {
			t.Errorf("Error is nil, when expected")
		}
		if err != ErrNotFound {
			t.Errorf("Error is not ErrNotFound, when expected. got: %v", err.Error())
		}
	})
	t.Run("Existing user", func(t *testing.T) {
		want := models.User{
			Name:     "name",
			Email:    "email",
			Password: "password",
			Role:     "role",
		}
		got, _ := userRepository.Save(&want)
		err := userService.DeleteUser(got.ID)
		if err != nil {
			t.Errorf("Error when delete user, when not expected. Error: %v", err)
		}
	})
}

func TestDecodeToken(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	cfg := config.Config{
		Db:  config.Db{},
		App: config.App{ApiSecret: "secret"},
	}
	userRepository, err := repository.NewUserRepository(db)
	if err != nil {
		t.Errorf("Error when create new user repository, when not expected. Error: %v", err)
	}
	userService := NewUserService(userRepository, &cfg)
	t.Run("Invalid token", func(t *testing.T) {
		_, err := userService.DecodeToken("invalid")
		if err == nil {
			t.Errorf("Error is nil, when expected")
		}
	})
	t.Run("Valid token", func(t *testing.T) {
		want := models.User{
			Name:     "name",
			Email:    "email",
			Password: "password",
			Role:     "role",
		}
		got, _ := userRepository.Save(&want)
		token, err := token.GenerateToken(got.ID, &cfg)
		if err != nil {
			t.Errorf("Error when generate token, when not expected. Error: %v", err)
		}
		_, err = userService.DecodeToken(token)
		if err != nil {
			t.Errorf("Error when decode token, when not expected. Error: %v", err)
		}
	})
	t.Run("Invalid user id", func(t *testing.T) {
		want := models.User{
			Name:     "name",
			Email:    "email",
			Password: "password",
			Role:     "role",
		}
		got, _ := userRepository.Save(&want)
		token, err := token.GenerateToken(got.ID+1, &cfg)
		if err != nil {
			t.Errorf("Error when generate token, when not expected. Error: %v", err)
		}
		_, err = userService.DecodeToken(token)
		if err == nil {
			t.Errorf("Error is nil, when expected")
		}
		if err != ErrNotFound {
			t.Errorf("Error is not ErrNotFound, when expected")
		}
	})
}

func compareUser(t *testing.T, got schemas.User, want models.User) {
	t.Helper()
	if got.Name != want.Name {
		t.Errorf("Name is not the same, got: %v, want: %v", got.Name, want.Name)
	}
	if got.Email != want.Email {
		t.Errorf("Email is not the same, got: %v, want: %v", got.Email, want.Email)
	}
	if got.Role != want.Role {
		t.Errorf("Role is not the same, got: %v, want: %v", got.Role, want.Role)
	}
}
