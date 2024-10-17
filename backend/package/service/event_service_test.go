package service

import (
	"testing"

	"github.com/HermanPlay/web-app-backend/package/domain/models"
	"github.com/HermanPlay/web-app-backend/package/domain/schemas"
	"github.com/HermanPlay/web-app-backend/package/repository"
	"github.com/HermanPlay/web-app-backend/package/utils"
)

func TestGetAllEvent(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	eventRepository, err := repository.NewEventRepository(db)
	if err != nil {
		t.Errorf("Error when create new event repository, when not expected. Error: %v", err)
	}

	userRepository, err := repository.NewUserRepository(db)
	if err != nil {
		t.Errorf("Error when create new user repository, when not expected. Error: %v", err)
	}
	user, err := userRepository.Save(&models.User{Name: "name", Email: "email", Password: "password", Role: "user"})
	if err != nil {
		t.Errorf("Error when save user, when not expected. Error: %v", err)
	}
	eventService := NewEventService(eventRepository)
	t.Run("Empty events", func(t *testing.T) {
		events, err := eventService.GetAllEvent()
		if err != nil {
			t.Errorf("Error when get all events, when not expected. Error: %v", err)
		}
		if len(events) != 0 {
			t.Errorf("Events is not empty, when expected")
		}
	})
	t.Run("One event", func(t *testing.T) {
		want := models.Event{
			Title:            "title",
			ShortDescription: "short description",
			Description:      "description",
			Location:         "location",
			Date:             "date",
			Time:             "time",
			CreatedBy:        user.ID,
		}
		_, err := eventRepository.Save(&want)
		if err != nil {
			t.Errorf("Error when save event, when not expected. Error: %v", err)
		}
		events, err := eventService.GetAllEvent()
		if err != nil {
			t.Errorf("Error when get all events, when not expected. Error: %v", err)
		}
		if len(events) != 1 {
			t.Errorf("Events is not empty, when expected")
		}
		compareEvent(t, *events[0], want)
	})
}

func TestGetEventById(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	eventRepository, err := repository.NewEventRepository(db)
	if err != nil {
		t.Errorf("Error when create new event repository, when not expected. Error: %v", err)
	}

	userRepository, err := repository.NewUserRepository(db)
	if err != nil {
		t.Errorf("Error when create new user repository, when not expected. Error: %v", err)
	}
	user, err := userRepository.Save(&models.User{Name: "name", Email: "email", Password: "password", Role: "user"})
	if err != nil {
		t.Errorf("Error when save user, when not expected. Error: %v", err)
	}
	eventService := NewEventService(eventRepository)
	t.Run("Invalid id", func(t *testing.T) {
		event, err := eventService.GetEventByID(1)
		if err == nil {
			t.Errorf("Error is nil, when expected")
		}
		if event != nil {
			t.Errorf("Event is not empty, when expected")
		}
	})
	t.Run("Valid id", func(t *testing.T) {
		want := models.Event{
			Title:            "title",
			ShortDescription: "short description",
			Description:      "description",
			Location:         "location",
			Date:             "date",
			Time:             "time",
			CreatedBy:        user.ID,
		}
		savedEvent, err := eventRepository.Save(&want)
		if err != nil {
			t.Errorf("Error when save event, when not expected. Error: %v", err)
		}
		event, err := eventService.GetEventByID(savedEvent.ID)
		if err != nil {
			t.Errorf("Error when get event by id, when not expected. Error: %v", err)
		}
		compareEvent(t, *event, want)
	})
}

func TestCreateEvent(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	eventRepository, err := repository.NewEventRepository(db)
	if err != nil {
		t.Errorf("Error when create new event repository, when not expected. Error: %v", err)
	}
	eventService := NewEventService(eventRepository)
	userRepository, err := repository.NewUserRepository(db)
	if err != nil {
		t.Errorf("Error when create new user repository, when not expected. Error: %v", err)
	}
	user, err := userRepository.Save(&models.User{Name: "name", Email: "email", Password: "password", Role: "user"})
	if err != nil {
		t.Errorf("Error when save user, when not expected. Error: %v", err)
	}
	t.Run("Create event", func(t *testing.T) {
		want := schemas.EventInput{
			Title:            "title",
			ShortDescription: "short description",
			Description:      "description",
			Location:         "location",
			Date:             "date",
			Time:             "time",
		}
		event, err := eventService.CreateEvent(&want, user.ID)
		if err != nil {
			t.Errorf("Error when create event, when not expected. Error: %v", err)
		}
		compareEvent(t, *event, models.Event{Title: want.Title, Description: want.Description, Location: want.Location, Date: want.Date, Time: want.Time, CreatedBy: user.ID})
	})
}

func TestUpdateEvent(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	eventRepository, err := repository.NewEventRepository(db)
	if err != nil {
		t.Errorf("Error when create new event repository, when not expected. Error: %v", err)
	}
	eventService := NewEventService(eventRepository)
	userRepository, err := repository.NewUserRepository(db)
	if err != nil {
		t.Errorf("Error when create new user repository, when not expected. Error: %v", err)
	}
	user, err := userRepository.Save(&models.User{Name: "name", Email: "email", Password: "password", Role: "user"})
	t.Run("Update event", func(t *testing.T) {
		want := models.Event{
			Title:            "title",
			ShortDescription: "short description",
			Description:      "description",
			Location:         "location",
			Date:             "date",
			Time:             "time",
			CreatedBy:        user.ID,
		}
		_, err := eventRepository.Save(&want)
		if err != nil {
			t.Errorf("Error when save event, when not expected. Error: %v", err)
		}
		update := schemas.EventUpdate{
			Title:            "new title",
			ShortDescription: "new short description",
			Description:      "new description",
			Location:         "new location",
			Date:             "new date",
			Time:             "new time",
		}
		event, err := eventService.UpdateEvent(&update, user.ID)
		if err != nil {
			t.Errorf("Error when update event, when not expected. Error: %v", err)
		}
		compareEvent(t, *event, models.Event{Title: update.Title, Description: update.Description, Location: update.Location, Date: update.Date, Time: update.Time, CreatedBy: user.ID})
	})
}

func TestDeleteEvent(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	eventRepository, err := repository.NewEventRepository(db)
	if err != nil {
		t.Errorf("Error when create new event repository, when not expected. Error: %v", err)
	}
	eventService := NewEventService(eventRepository)
	userRepository, err := repository.NewUserRepository(db)
	if err != nil {
		t.Errorf("Error when create new user repository, when not expected. Error: %v", err)
	}
	user, err := userRepository.Save(&models.User{Name: "name", Email: "email", Password: "password", Role: "user"})
	t.Run("Delete event", func(t *testing.T) {
		want := models.Event{
			Title:            "title",
			ShortDescription: "short description",
			Description:      "description",
			Location:         "location",
			Date:             "date",
			Time:             "time",
			CreatedBy:        user.ID,
		}
		savedEvent, err := eventRepository.Save(&want)
		if err != nil {
			t.Errorf("Error when save event, when not expected. Error: %v", err)
		}
		err = eventService.DeleteEvent(savedEvent.ID)
		if err != nil {
			t.Errorf("Error when delete event, when not expected. Error: %v", err)
		}
		_, err = eventRepository.GetByID(savedEvent.ID)
		if err == nil {
			t.Errorf("Error is nil, when expected")
		}
	})
}
func TestBookEvent(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	eventRepository, err := repository.NewEventRepository(db)
	if err != nil {
		t.Errorf("Error when create new event repository, when not expected. Error: %v", err)
	}
	eventService := NewEventService(eventRepository)
	userRepository, err := repository.NewUserRepository(db)
	if err != nil {
		t.Errorf("Error when create new user repository, when not expected. Error: %v", err)
	}
	user, err := userRepository.Save(&models.User{Name: "name", Email: "email", Password: "password", Role: "user"})
	if err != nil {
		t.Errorf("Error when save user, when not expected. Error: %v", err)
	}
	t.Run("Book event", func(t *testing.T) {
		want := models.Event{
			Title:            "title",
			ShortDescription: "short description",
			Description:      "description",
			Location:         "location",
			Date:             "date",
			Time:             "time",
			CreatedBy:        user.ID,
		}
		savedEvent, err := eventRepository.Save(&want)
		if err != nil {
			t.Errorf("Error when save event, when not expected. Error: %v", err)
		}
		err = eventService.BookEvent(savedEvent.ID, user.ID)
		if err != nil {
			t.Errorf("Error when book event, when not expected. Error: %v", err)
		}
	})
}

func TestGetFeaturedEvents(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	eventRepository, err := repository.NewEventRepository(db)
	if err != nil {
		t.Errorf("Error when create new event repository, when not expected. Error: %v", err)
	}
	eventService := NewEventService(eventRepository)
	userRepository, err := repository.NewUserRepository(db)
	if err != nil {
		t.Errorf("Error when create new user repository, when not expected. Error: %v", err)
	}
	user, err := userRepository.Save(&models.User{Name: "name", Email: "email", Password: "password", Role: "user"})
	if err != nil {
		t.Errorf("Error when save user, when not expected. Error: %v", err)
	}
	t.Run("Get featured events", func(t *testing.T) {
		want := models.Event{
			Title:            "title",
			ShortDescription: "short description",
			Description:      "description",
			Location:         "location",
			Date:             "date",
			Time:             "time",
			CreatedBy:        user.ID,
			IsFeatured:       true,
		}
		_, err := eventRepository.Save(&want)
		if err != nil {
			t.Errorf("Error when save event, when not expected. Error: %v", err)
		}
		events, err := eventService.GetFeaturedEvents()
		if err != nil {
			t.Errorf("Error when get featured events, when not expected. Error: %v", err)
		}
		if len(events) != 1 {
			t.Errorf("Events is not empty, when expected")
		}
		compareEvent(t, *events[0], want)
	})
}

func TestGetMyEvents(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	eventRepository, err := repository.NewEventRepository(db)
	if err != nil {
		t.Errorf("Error when create new event repository, when not expected. Error: %v", err)
	}
	eventService := NewEventService(eventRepository)
	userRepository, err := repository.NewUserRepository(db)
	if err != nil {
		t.Errorf("Error when create new user repository, when not expected. Error: %v", err)
	}
	user, err := userRepository.Save(&models.User{Name: "name", Email: "email", Password: "password", Role: "user"})
	if err != nil {
		t.Errorf("Error when save user, when not expected. Error: %v", err)
	}
	t.Run("Get my events", func(t *testing.T) {
		want := models.Event{
			Title:       "title",
			Description: "description",
			Location:    "location",
			Date:        "date",
			Time:        "time",
			CreatedBy:   user.ID,
		}
		_, err := eventRepository.Save(&want)
		if err != nil {
			t.Errorf("Error when save event, when not expected. Error: %v", err)
		}
		events, err := eventService.GetMyEvents(user.ID)
		if err != nil {
			t.Errorf("Error when get my events, when not expected. Error: %v", err)
		}
		if len(events) != 1 {
			t.Errorf("Events is not empty, when expected")
		}
		compareEvent(t, *events[0], want)
	})
}

func compareEvent(t *testing.T, got schemas.Event, want models.Event) {
	t.Helper()
	if got.Title != want.Title {
		t.Errorf("Title is not the same, got: %v, want: %v", got.Title, want.Title)
	}
	if got.Description != want.Description {
		t.Errorf("Description is not the same, got: %v, want: %v", got.Description, want.Description)
	}
	if got.Location != want.Location {
		t.Errorf("Location is not the same, got: %v, want: %v", got.Location, want.Location)
	}
	if got.Date != want.Date {
		t.Errorf("Date is not the same, got: %v, want: %v", got.Date, want.Date)
	}
	if got.Time != want.Time {
		t.Errorf("Time is not the same, got: %v, want: %v", got.Time, want.Time)
	}
	if got.CreatedBy != want.CreatedBy {
		t.Errorf("CreatedBy is not the same, got: %v, want: %v", got.CreatedBy, want.CreatedBy)
	}
}
