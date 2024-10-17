package repository

import (
	"testing"
	"time"

	"github.com/HermanPlay/web-app-backend/package/domain/models"
	"github.com/HermanPlay/web-app-backend/package/utils"
	"gorm.io/gorm"
)

func TestNewEventRepository(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	eventRepo, err := NewEventRepository(db)
	if err != nil {
		t.Errorf("Error when create new event repository, when not expected. Error: %v", err)
	}
	if eventRepo == nil {
		t.Error("Event repository is nil")
	}
}

func TestGetAll(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	eventRepo, _ := NewEventRepository(db)
	events, err := eventRepo.GetAll()
	if err != nil {
		t.Errorf("Error when get all events, when not expected. Error: %v", err)
	}
	if len(events) != 0 {
		t.Errorf("Events is not empty, when expected")
	}
}

func TestGetByID(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	eventRepo, _ := NewEventRepository(db)
	event, err := eventRepo.GetByID(1)
	if err == nil {
		t.Errorf("Error is nil, when expected")
	}
	if event.ID != 0 {
		t.Errorf("Event is not empty, when expected")
	}
}

func TestSaveEvent(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	eventRepo, _ := NewEventRepository(db)
	want := models.Event{
		Title:            "event",
		ShortDescription: "short description",
		Description:      "description",
		Location:         "location",
		Date:             time.Now().Format("2006-01-02"),
		Time:             time.Now().Format("15:04"),
		CreatedBy:        1,
	}
	createUser(db)
	got, err := eventRepo.Save(&want)
	if err != nil {
		t.Errorf("Ereor when save event, when not expected. Error: %v", err)
		return
	}
	compareEvent(t, got, want)
}

func TestUpdateEvent(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	eventRepo, _ := NewEventRepository(db)
	want := models.Event{
		Title:            "event",
		ShortDescription: "short description",
		Description:      "description",
		Location:         "location",
		Date:             time.Now().Format("2006-01-02"),
		Time:             time.Now().Format("15:04"),
		CreatedBy:        1,
	}
	createUser(db)
	got, err := eventRepo.Save(&want)
	if err != nil {
		t.Errorf("Error when save event, when not expected. Error: %v", err)
	}
	compareEvent(t, got, want)
	want.Title = "event2"
	want.ShortDescription = "short description2"
	want.Description = "description2"
	want.Location = "location2"
	want.Date = time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	want.Time = time.Now().Add(1 * time.Hour).Format("15:04")
	got, err = eventRepo.Update(&want)
	if err != nil {
		t.Errorf("Error when update event, when not expected. Error: %v", err)
	}
	compareEvent(t, got, want)
}

func TestDeleteEvent(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	eventRepo, _ := NewEventRepository(db)
	want := models.Event{
		Title:            "event",
		ShortDescription: "short description",
		Description:      "description",
		Location:         "location",
		Date:             time.Now().Format("2006-01-02"),
		Time:             time.Now().Format("15:04"),
		CreatedBy:        1,
	}
	createUser(db)
	got, err := eventRepo.Save(&want)
	if err != nil {
		t.Errorf("Error when save event, when not expected. Error: %v", err)
	}
	compareEvent(t, got, want)
	err = eventRepo.Delete(got.ID)
	if err != nil {
		t.Errorf("Error when delete event, when not expected. Error: %v", err)
	}
}

func TestGetFeaturedEvents(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	eventRepo, _ := NewEventRepository(db)
	events, err := eventRepo.GetFeaturedEvents()
	if err != nil {
		t.Errorf("Error when get featured events, when not expected. Error: %v", err)
	}
	if len(events) != 0 {
		t.Errorf("Events is not empty, when expected")
	}
}

func TestGetCreatedEvents(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	eventRepo, _ := NewEventRepository(db)
	events, err := eventRepo.GetCreatedEvents(1)
	if err != nil {
		t.Errorf("Error when get created events, when not expected. Error: %v", err)
	}
	if len(events) != 0 {
		t.Errorf("Events is not empty, when expected")
	}
}

func TestGetMyEvents(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	eventRepo, _ := NewEventRepository(db)
	events, err := eventRepo.GetMyEvents(1)
	if err != nil {
		t.Errorf("Error when get my events, when not expected. Error: %v", err)
	}
	if len(events) != 0 {
		t.Errorf("Events is not empty, when expected")
	}
}

func TestBookEvent(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	eventRepo, _ := NewEventRepository(db)
	event := models.Event{
		Title:            "event",
		ShortDescription: "short description",
		Description:      "description",
		Location:         "location",
		Date:             time.Now().Format("2006-01-02"),
		Time:             time.Now().Format("15:04"),
		CreatedBy:        1,
	}
	createUser(db)
	eventRepo.Save(&event)
	err := eventRepo.BookEvent(1, 1)
	if err != nil {
		t.Errorf("Error when book event, when not expected. Error: %v", err)
	}
}

func TestGetBooking(t *testing.T) {
	db := utils.ConnectToTestDatabase()
	eventRepo, _ := NewEventRepository(db)
	event := models.Event{
		Title:            "event",
		ShortDescription: "short description",
		Description:      "description",
		Location:         "location",
		Date:             time.Now().Format("2006-01-02"),
		Time:             time.Now().Format("15:04"),
		CreatedBy:        1,
	}
	createUser(db)
	eventRepo.Save(&event)
	eventRepo.BookEvent(1, 1)
	booking, err := eventRepo.GetBooking(1, 1)
	if err != nil {
		t.Errorf("Error when get booking, when not expected. Error: %v", err)
	}
	if booking.EventID != 1 {
		t.Errorf("Booking for event is not same, got: %d, want: %d", booking.EventID, 1)
	}

	if booking.UserID != 1 {
		t.Errorf("Booking for user is not same, got: %d, want: %d", booking.UserID, 1)
	}

}

func createUser(db *gorm.DB) {
	user := models.User{
		Email:    "email",
		Password: "password",
		Role:     models.UserRole,
	}
	userRepo, _ := NewUserRepository(db)
	userRepo.Save(&user)
}

func compareEvent(t *testing.T, got, want models.Event) {
	t.Helper()
	if got.Title != want.Title {
		t.Errorf("Title is not same, got: %s, want: %s", got.Title, want.Title)
	}
	if got.ShortDescription != want.ShortDescription {
		t.Errorf("ShortDescription is not same, got: %s, want: %s", got.ShortDescription, want.ShortDescription)
	}
	if got.Description != want.Description {
		t.Errorf("Description is not same, got: %s, want: %s", got.Description, want.Description)
	}
	if got.Location != want.Location {
		t.Errorf("Location is not same, got: %s, want: %s", got.Location, want.Location)
	}
	if got.Date != want.Date {
		t.Errorf("Date is not same, got: %s, want: %s", got.Date, want.Date)
	}
	if got.Time != want.Time {
		t.Errorf("Time is not same, got: %s, want: %s", got.Time, want.Time)
	}
	if got.CreatedBy != want.CreatedBy {
		t.Errorf("CreatedBy is not same, got: %d, want: %d", got.CreatedBy, want.CreatedBy)
	}
}
