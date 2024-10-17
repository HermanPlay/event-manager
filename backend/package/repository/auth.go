package repository

import (
	"log"

	"github.com/HermanPlay/web-app-backend/internal/api/http/util/token"
	"github.com/HermanPlay/web-app-backend/internal/config"
	"github.com/HermanPlay/web-app-backend/package/domain/models"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository interface {
	LoginUser(email, password string) (string, error)
}

type AuthRepositoryImpl struct {
	db  *gorm.DB
	cfg *config.Config
}

func verifyPassword(inputPassword, validPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(validPassword), []byte(inputPassword))
}

func (a AuthRepositoryImpl) LoginUser(email, password string) (string, error) {
	var user models.User
	err := a.db.Model(&user).Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", err
		}
		logrus.Error("Error getting user from db, err: ", err.Error())
		return "", err
	}

	log.Printf("user: %v", user)

	logrus.Printf("user: %v", user)

	err = verifyPassword(password, user.Password)
	if err != nil {
		return "", err
	}

	token, err := token.GenerateToken(user.ID, a.cfg)
	if err != nil {
		return "", err
	}
	return token, nil
}

func NewAuthRepository(db *gorm.DB, cfg *config.Config) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{
		db:  db,
		cfg: cfg,
	}
}
