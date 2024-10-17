package routes

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/HermanPlay/web-app-backend/internal/api/http/constant"
	"github.com/HermanPlay/web-app-backend/internal/api/http/util"
	"github.com/HermanPlay/web-app-backend/package/domain/schemas"
	"github.com/HermanPlay/web-app-backend/package/service"
	"github.com/gin-gonic/gin"
)

type UserRoute interface {
	GetAllUserData(c *gin.Context)
	AddUserData(c *gin.Context)
	GetUserById(c *gin.Context)
	UpdateUserData(c *gin.Context)
	DeleteUser(c *gin.Context)
	DecodeToken(c *gin.Context)
}

type UserRouteImpl struct {
	service service.UserService
}

func (u UserRouteImpl) GetAllUserData(c *gin.Context) {
	data, err := u.service.GetAllUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, util.BuildResponse(constant.InvalidRequest, "Error when getting data"))
		return
	}
	c.JSON(http.StatusOK, util.BuildResponse(constant.Success, data))
}

func (u UserRouteImpl) AddUserData(c *gin.Context) {
	var data schemas.UserInput
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, util.BuildResponse(constant.InvalidRequest, "Invalid input. Check your input types."))
		return
	}

	user, err := u.service.AddUserData(data)
	if err != nil {
		if err == service.ErrAlreadyExists {
			c.JSON(http.StatusConflict, util.BuildResponse(constant.AlreadyExists, "User with given email already exists"))
			return
		}
		if err == service.ErrInvalidInput {
			c.JSON(http.StatusBadRequest, util.BuildResponse(constant.InvalidRequest, "Empty input. Check your input values."))
			return
		}
		c.JSON(http.StatusInternalServerError, util.BuildResponse(constant.UnknownError, "Unknown internal server error"))
		return
	}
	c.JSON(http.StatusCreated, util.BuildResponse(constant.Success, user))
}

func (u UserRouteImpl) GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.BuildResponse(constant.InvalidRequest, "Invalid id supplied"))
	}

	data, err := u.service.GetUserById(id)
	if err != nil {
		if err == service.ErrNotFound {
			c.JSON(http.StatusNotFound, util.BuildResponse(constant.NotFound, "User not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, util.BuildResponse(constant.UnknownError, "Unknown internal server error"))
		return
	}
	c.JSON(http.StatusOK, util.BuildResponse(constant.Success, data))
}

func (u UserRouteImpl) UpdateUserData(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.BuildResponse(constant.InvalidRequest, "Invalid id supplied"))
	}
	var user schemas.UserUpdate
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, util.BuildResponse(constant.InvalidRequest, "Invalid input. Check your input types."))
		return
	}

	data, err := u.service.UpdateUserData(user, id)
	if err != nil {
		if err == service.ErrNotFound {
			c.JSON(http.StatusNotFound, util.BuildResponse(constant.NotFound, "User not found"))
			return
		}
		if err == service.ErrAlreadyExists {
			c.JSON(http.StatusConflict, util.BuildResponse(constant.AlreadyExists, "User with given email already exists"))
			return
		}
		c.JSON(http.StatusInternalServerError, util.BuildResponse(constant.UnknownError, "Unknown internal server error"))
		return
	}

	c.JSON(http.StatusOK, util.BuildResponse(constant.Success, data))

}

func (u UserRouteImpl) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.BuildResponse(constant.InvalidRequest, "Invalid id supplied"))
		return
	}

	err = u.service.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.BuildResponse(constant.InvalidRequest, "Unkonwn internal server error"))
		return
	}
}

func (u UserRouteImpl) DecodeToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	token := strings.Split(authHeader, "Bearer ")[1]
	user, err := u.service.DecodeToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.BuildResponse(constant.InvalidRequest, "Could not decode token. "+err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.BuildResponse(constant.Success, user))
}

func NewUserRoute(userService service.UserService) UserRoute {
	return &UserRouteImpl{
		service: userService,
	}
}
