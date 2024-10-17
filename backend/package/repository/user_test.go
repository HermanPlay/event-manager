package repository

import (
	"testing"

	"github.com/HermanPlay/web-app-backend/package/domain/models"
	"github.com/HermanPlay/web-app-backend/package/utils"
)

func TestNewUserRepository(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	userRepositoryImpl, err := NewUserRepository(db)
	if err != nil {
		t.Errorf("Error when create new user repository, when not expected. Error: %v", err)
	}
	if userRepositoryImpl == nil {
		t.Errorf("User repository is nil, when not expected")
	}
}

func TestFindAllUser(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	userRepositoryImpl, _ := NewUserRepository(db)
	users, err := userRepositoryImpl.FindAllUser()
	if err != nil {
		t.Errorf("Error when find all user, when not expected. Error: %v", err)
	}
	if len(users) != 0 {
		t.Errorf("Users is not empty, when expected")
	}
}

func TestSave(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	userRepositoryImpl, _ := NewUserRepository(db)
	want := models.User{
		Email:    "email@email.com",
		Password: "password",
		Role:     models.UserRole,
	}
	got, err := userRepositoryImpl.Save(&want)
	if err != nil {
		t.Errorf("Error when save user, when not expected. Error: %v", err)
	}
	compareUser(t, got, want)
	want.Email = "email2@email.com"
	want.Role = models.ManagerRole
	got, err = userRepositoryImpl.Save(&want)
	if err != nil {
		t.Errorf("Error when save user, when not expected. Error: %v", err)
	}
	compareUser(t, got, want)
	want.Email = "email3@email.com"
	want.Role = models.AdminRole
	got, err = userRepositoryImpl.Save(&want)
	if err != nil {
		t.Errorf("Error when save user, when not expected. Error: %v", err)
	}
	compareUser(t, got, want)
}

func TestFindUserById(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	userRepositoryImpl, _ := NewUserRepository(db)
	want := models.User{
		Email:    "email5@email.com",
		Password: "password",
		Role:     models.UserRole,
	}
	got, _ := userRepositoryImpl.Save(&want)
	user, err := userRepositoryImpl.FindUserById(got.ID)
	if err != nil {
		t.Errorf("Error when find user by id, when not expected. Error: %v", err)
	}
	compareUser(t, user, want)
}

func TestDeleteUserById(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	userRepositoryImpl, _ := NewUserRepository(db)
	want := models.User{
		Email:    "email6@email.com",
		Password: "password",
		Role:     models.UserRole,
	}
	got, _ := userRepositoryImpl.Save(&want)
	err := userRepositoryImpl.DeleteUserById(got.ID)
	if err != nil {
		t.Errorf("Error when delete user by id, when not expected. Error: %v", err)
	}
	user, err := userRepositoryImpl.FindUserById(got.ID)
	if err == nil {
		t.Errorf("User is not deleted, when expected. User: %v", user)
	}
}

func TestCheckUserExist(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	userRepositoryImpl, _ := NewUserRepository(db)
	exist, err := userRepositoryImpl.CheckUserExist("email7@email.com")
	if err != nil {
		t.Errorf("Error when check user exist, when not expected. Error: %v", err)
	}
	if exist {
		t.Errorf("User exist, when not expected")
	}

	want := models.User{
		Email:    "email7@email.com",
		Password: "password",
		Role:     models.UserRole,
	}
	got, _ := userRepositoryImpl.Save(&want)
	exist, err = userRepositoryImpl.CheckUserExist(got.Email)
	if err != nil {
		t.Errorf("Error when check user exist, when not expected. Error: %v", err)
	}
	if !exist {
		t.Errorf("User is not exist, when expected")
	}
}

func TestUpdate(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	userRepositoryImpl, _ := NewUserRepository(db)
	want := models.User{
		Email:    "email8@email.com",
		Password: "password",
		Role:     models.UserRole,
	}
	userRepositoryImpl.Save(&want)
	want.Email = "newemail8@email.com"
	want.Role = models.ManagerRole
	got, err := userRepositoryImpl.Update(&want)
	if err != nil {
		t.Errorf("Error when update user, when not expected. Error: %v", err)
	}
	compareUser(t, got, want)
}

func TestGetUserByEmail(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	userRepositoryImpl, _ := NewUserRepository(db)
	got, err := userRepositoryImpl.GetUserByEmail("email9@email.com")
	if err == nil {
		t.Errorf("User is exist, when not expected. User: %v", got)
	}

	want := models.User{
		Email:    "email9@email.com",
		Password: "password",
		Role:     models.UserRole,
	}
	got, _ = userRepositoryImpl.Save(&want)
	user, err := userRepositoryImpl.GetUserByEmail(got.Email)
	if err != nil {
		t.Errorf("Error when get user by email, when not expected. Error: %v", err)
	}
	compareUser(t, user, want)
}

func compareUser(t *testing.T, got, want models.User) {
	if got.Email != want.Email {
		t.Errorf("Email is not same, got: %s, want: %s", got.Email, want.Email)
	}
	if got.Password != want.Password {
		t.Errorf("Password is not same, got: %s, want: %s", got.Password, want.Password)
	}
	if got.Role != want.Role {
		t.Errorf("Role is not same, got: %s, want: %s", got.Role, want.Role)
	}
}
