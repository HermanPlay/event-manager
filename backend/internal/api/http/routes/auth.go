package routes

import (
	"net/http"
	"time"

	"github.com/HermanPlay/web-app-backend/internal/api/http/constant"
	"github.com/HermanPlay/web-app-backend/internal/api/http/util"
	"github.com/HermanPlay/web-app-backend/package/domain/schemas"
	"github.com/HermanPlay/web-app-backend/package/service"
	"github.com/gin-gonic/gin"
)

type AuthRoute interface {
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
	ResetPassword(c *gin.Context)
}

type AuthRouteImpl struct {
	service service.AuthService
}

func (a AuthRouteImpl) RegisterUser(c *gin.Context) {
	var userRegister schemas.UserRegister
	if err := c.ShouldBindJSON(&userRegister); err != nil {
		c.JSON(http.StatusBadRequest, util.BuildResponse(constant.InvalidRequest, "Invalid request data"))
		return
	}

	data, err := a.service.RegisterUser(userRegister)
	if err != nil {
		if err == service.ErrAlreadyExists {
			c.JSON(http.StatusConflict, util.BuildResponse(constant.AlreadyExists, "User with given email already exists"))
			return
		}
		c.JSON(http.StatusInternalServerError, util.BuildResponse(constant.UnknownError, "Unknown internal server error"))
		return
	}
	c.JSON(http.StatusOK, util.BuildResponse(constant.Success, data))
}

func (a AuthRouteImpl) LoginUser(c *gin.Context) {
	var userLogin schemas.UserLogin
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, util.BuildResponse(constant.InvalidRequest, "Invalid request data"))
		return
	}

	token, err := a.service.LoginUser(userLogin)
	if err != nil {
		if err == service.ErrNotFound {
			c.JSON(http.StatusNotFound, util.BuildResponse(constant.NotFound, "User not found"))
			return
		}
		if err == service.ErrInvalidPassword {
			c.JSON(http.StatusUnauthorized, util.BuildResponse(constant.InvalidCredentials, "Invalid password"))
			return
		}
		c.JSON(http.StatusInternalServerError, util.BuildResponse(constant.UnknownError, "Unknown internal server error"))
		return
	}
	c.SetCookie("token", token, int((constant.TokenHourLifespan * time.Hour).Seconds()), "/", "localhost", false, true)
	c.JSON(http.StatusOK, util.BuildResponse(constant.Success, token))
}

func (a AuthRouteImpl) ResetPassword(c *gin.Context) {
	var userResetPassword schemas.UserResetPassword
	if err := c.ShouldBindJSON(&userResetPassword); err != nil {
		c.JSON(http.StatusBadRequest, util.BuildResponse(constant.InvalidRequest, "Invalid request data"))
		return
	}

	newPassword, err := a.service.ResetPassword(userResetPassword)
	if err != nil {
		if err == service.ErrNotFound {
			c.JSON(http.StatusNotFound, util.BuildResponse(constant.NotFound, "User not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, util.BuildResponse(constant.UnknownError, "Unknown internal server error"))
		return
	}
	c.JSON(http.StatusOK, util.BuildResponse(constant.Success, newPassword))
}

func NewAuthRoute(service service.AuthService) *AuthRouteImpl {
	return &AuthRouteImpl{
		service: service,
	}
}
