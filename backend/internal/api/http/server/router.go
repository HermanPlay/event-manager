package server

import (
	"time"

	"github.com/HermanPlay/web-app-backend/internal/api/http"
	"github.com/HermanPlay/web-app-backend/internal/api/http/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(init *http.Initialization) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	cors.Default()
	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,

		MaxAge: 12 * time.Hour,
	}
	api.Use(middleware.CORSMiddleware(corsConfig))
	api.OPTIONS("/*path", cors.New(corsConfig))
	{
		dev := api.Group("/dev")
		dev.GET("/status", init.DevRoute.HealthCheck)

		auth := api.Group("/auth")
		auth.POST("/register", init.AuthRoute.RegisterUser)
		auth.POST("/login", init.AuthRoute.LoginUser)
		auth.POST("/reset", init.AuthRoute.ResetPassword)

		user := api.Group("/user")
		user.Use(middleware.JwtAuthMiddleware(init.Cfg))
		user.GET("", init.UserRoute.GetAllUserData)
		user.POST("", init.UserRoute.AddUserData)
		user.GET("/:userID", init.UserRoute.GetUserById)
		user.PATCH("/:userID", init.UserRoute.UpdateUserData)
		user.DELETE("/:userID", init.UserRoute.DeleteUser)
		user.GET("/decode", init.UserRoute.DecodeToken)

		// Can be accessed without authentication
		api.GET("/event/featured", init.EventRoute.GetFeaturedEvents)
		event := api.Group("/event")
		event.Use(middleware.JwtAuthMiddleware(init.Cfg))
		event.GET("", init.EventRoute.GetAllEvent)
		event.GET("/my/:userID", init.EventRoute.GetMyEvents)
		event.POST("/book/:eventID", init.EventRoute.BookEvent)
		event.GET("/:eventID", init.EventRoute.GetEventById)
		event.POST("", init.EventRoute.CreateEvent)
		event.PATCH("/:eventID", init.EventRoute.UpdateEvent)
		event.DELETE("/:eventID", init.EventRoute.DeleteEvent)
	}

	return router
}
