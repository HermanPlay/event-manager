package routes

import (
	"net/http"

	"github.com/HermanPlay/web-app-backend/internal/api/http/constant"
	"github.com/HermanPlay/web-app-backend/internal/api/http/util"
	"github.com/gin-gonic/gin"
)

type DevRoute interface {
	HealthCheck(c *gin.Context)
}

type DevRouteImpl struct {
}

func (d DevRouteImpl) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, util.BuildResponse(constant.Success, map[string]string{"message": "ok"}))
}

func NewDevRoute() DevRoute {
	return DevRouteImpl{}
}
