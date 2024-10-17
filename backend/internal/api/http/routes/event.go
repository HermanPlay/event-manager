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

type EventRoute interface {
	GetAllEvent(c *gin.Context)
	GetEventById(c *gin.Context)
	CreateEvent(c *gin.Context)
	UpdateEvent(c *gin.Context)
	DeleteEvent(c *gin.Context)
	GetFeaturedEvents(c *gin.Context)
	GetMyEvents(c *gin.Context)
	BookEvent(c *gin.Context)
}

type EventRouteImpl struct {
	eventService service.EventService
	userService  service.UserService
}

func NewEventRoute(eventService service.EventService, userService service.UserService) EventRoute {
	return &EventRouteImpl{
		eventService: eventService,
		userService:  userService,
	}
}

func (e EventRouteImpl) GetAllEvent(c *gin.Context) {
	data, err := e.eventService.GetAllEvent()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error when getting data"})
		return
	}
	c.JSON(http.StatusOK, util.BuildResponse(constant.Success, data))
}

func (e EventRouteImpl) GetEventById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("eventID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id supplied"})
		return
	}

	data, err := e.eventService.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error when getting data"})
		return
	}

	c.JSON(http.StatusOK, util.BuildResponse(constant.Success, data))
}

func (e EventRouteImpl) CreateEvent(c *gin.Context) {
	var eventInput schemas.EventInput
	if err := c.ShouldBindJSON(&eventInput); err != nil {
		c.JSON(http.StatusBadRequest, util.BuildResponse(constant.InvalidRequest, "Could not parse input body! Verify your input."+err.Error()))
		return
	}

	authHeader := c.GetHeader("Authorization")
	token := strings.Split(authHeader, "Bearer ")[1]
	user, err := e.userService.DecodeToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, util.BuildResponse(constant.InvalidRequest, "Could not decode user token"))
	}
	userId := user.ID
	data, err := e.eventService.CreateEvent(&eventInput, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error when saving data"})
		return
	}

	c.JSON(http.StatusOK, util.BuildResponse(constant.Success, data))
}

func (e EventRouteImpl) UpdateEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("eventID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id supplied"})
		return
	}

	var eventUpdate schemas.EventUpdate
	if err := c.ShouldBindJSON(&eventUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data schema, verify data types"})
		return
	}

	data, err := e.eventService.UpdateEvent(&eventUpdate, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error when updating data"})
		return
	}

	c.JSON(http.StatusOK, util.BuildResponse(constant.Success, data))
}

func (e EventRouteImpl) DeleteEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("eventID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id supplied"})
		return
	}

	err = e.eventService.DeleteEvent(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error when deleting data"})
		return
	}

	c.JSON(http.StatusOK, util.BuildResponse(constant.Success, map[string]string{"message": "Event deleted"}))
}

func (e EventRouteImpl) GetFeaturedEvents(c *gin.Context) {
	data, err := e.eventService.GetFeaturedEvents()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error when getting data"})
		return
	}
	c.JSON(http.StatusOK, util.BuildResponse(constant.Success, data))
}

func (e EventRouteImpl) GetMyEvents(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id supplied"})
		return
	}

	data, err := e.eventService.GetMyEvents(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error when getting data"})
		return
	}
	c.JSON(http.StatusOK, util.BuildResponse(constant.Success, data))
}

func (e EventRouteImpl) BookEvent(c *gin.Context) {
	eventID, err := strconv.Atoi(c.Param("eventID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id supplied"})
		return
	}

	authHeader := c.GetHeader("Authorization")
	token := strings.Split(authHeader, "Bearer ")[1]
	user, err := e.userService.DecodeToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, util.BuildResponse(constant.InvalidRequest, "Could not decode user token"))
	}
	userID := user.ID

	err = e.eventService.BookEvent(eventID, userID)
	if err != nil {
		if err == service.ErrBookingExists {
			c.JSON(http.StatusConflict, util.BuildResponse(constant.AlreadyExists, "Event already booked"))
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error when booking event"})
		return
	}

	c.JSON(http.StatusOK, util.BuildResponse(constant.Success, []string{"Event booked successfully"}))
}
