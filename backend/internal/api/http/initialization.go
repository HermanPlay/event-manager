package http

import (
	"fmt"

	"github.com/HermanPlay/web-app-backend/internal/api/http/routes"
	"github.com/HermanPlay/web-app-backend/internal/config"
	"github.com/HermanPlay/web-app-backend/internal/database"
	"github.com/HermanPlay/web-app-backend/package/domain/models"
	"github.com/HermanPlay/web-app-backend/package/repository"
	"github.com/HermanPlay/web-app-backend/package/service"
)

type Initialization struct {
	Cfg             *config.Config
	DevRoute        routes.DevRoute
	UserRepository  repository.UserRepository
	UserService     service.UserService
	UserRoute       routes.UserRoute
	AuthRepository  repository.AuthRepository
	AuthService     service.AuthService
	AuthRoute       routes.AuthRoute
	EventRepository repository.EventRepository
	EventService    service.EventService
	EventRoute      routes.EventRoute
}

func NewInitialization(
	config *config.Config,
	devRoute routes.DevRoute,
	userRepo repository.UserRepository,
	userService service.UserService,
	UserRoute routes.UserRoute,
	authRepo repository.AuthRepository,
	authService service.AuthService,
	authRoute routes.AuthRoute,
	eventRepository repository.EventRepository,
	eventService service.EventService,
	eventRoute routes.EventRoute,
) *Initialization {
	return &Initialization{
		Cfg:             config,
		DevRoute:        devRoute,
		UserRepository:  userRepo,
		UserService:     userService,
		UserRoute:       UserRoute,
		AuthRepository:  authRepo,
		AuthService:     authService,
		AuthRoute:       authRoute,
		EventRepository: eventRepository,
		EventService:    eventService,
		EventRoute:      eventRoute,
	}
}

func Init(cfg *config.Config) *Initialization {
	db, err := database.NewPostgresDatabase(cfg)
	if err != nil {
		panic(err)
	}
	pgDb := db.Connect()
	devRouteImpl := routes.NewDevRoute()
	userRepositoryImpl, err := repository.NewUserRepository(pgDb)
	if err != nil {
		panic(err)
	}
	userServiceImpl := service.NewUserService(userRepositoryImpl, cfg)
	userRouteImpl := routes.NewUserRoute(userServiceImpl)
	authRepositoryImpl := repository.NewAuthRepository(pgDb, cfg)
	authServiceImpl := service.NewAuthService(authRepositoryImpl, userRepositoryImpl)
	authRouteImpl := routes.NewAuthRoute(authServiceImpl)
	eventRepositoryImpl, err := repository.NewEventRepository(pgDb)
	if err != nil {
		panic(err)
	}
	eventServiceImpl := service.NewEventService(eventRepositoryImpl)
	eventRouteImpl := routes.NewEventRoute(eventServiceImpl, userServiceImpl)
	initialization := NewInitialization(cfg, devRouteImpl, userRepositoryImpl, userServiceImpl, userRouteImpl, authRepositoryImpl, authServiceImpl, authRouteImpl, eventRepositoryImpl, eventServiceImpl, eventRouteImpl)

	var count int64
	pgDb.Model(&models.User{}).Count(&count)

	if count == 0 {
		fmt.Println("Seeding database...")

		// Seed users
		users := []models.User{
			{Name: "John Doe", Email: "johndoe@example.com", Password: "password123", Role: "admin"},
			{Name: "Jane Smith", Email: "janesmith@example.com", Password: "password123", Role: "user"},
			{Name: "Mike Johnson", Email: "mikejohnson@example.com", Password: "password123", Role: "user"},
			{Name: "Alice Brown", Email: "alicebrown@example.com", Password: "password123", Role: "user"},
		}
		pgDb.Create(&users)

		// Seed events
		events := []models.Event{
			{Title: "Tech Conference 2024", Description: "A conference for tech enthusiasts.", Location: "New York", Date: "2024-11-15", Time: "09:00 AM", IsFeatured: true, CreatedBy: 1, ShortDescription: "Tech event for 2024"},
			{Title: "Music Festival", Description: "An outdoor music festival.", Location: "Los Angeles", Date: "2024-12-05", Time: "04:00 PM", IsFeatured: true, CreatedBy: 2, ShortDescription: "Enjoy live music all day"},
		}
		pgDb.Create(&events)

		fmt.Println("Database seeded successfully.")
	} else {
		fmt.Println("Database already seeded.")
	}
	return initialization
}
