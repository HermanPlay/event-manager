package repository

import (
	"github.com/HermanPlay/web-app-backend/package/domain/models"
	"gorm.io/gorm"
)

type EventRepository interface {
	GetAll() ([]models.Event, error)
	GetByID(id int) (models.Event, error)
	Save(event *models.Event) (models.Event, error)
	Update(event *models.Event) (models.Event, error)
	Delete(id int) error
	GetFeaturedEvents() ([]models.Event, error)
	GetMyEvents(userId int) ([]models.Event, error)
	GetCreatedEvents(userId int) ([]models.Event, error)
	BookEvent(eventID int, userID int) error
	GetBooking(eventID int, userID int) (models.EventUser, error)
}

type EventRepositoryImpl struct {
	db *gorm.DB
}

func (e EventRepositoryImpl) GetAll() ([]models.Event, error) {
	var events []models.Event
	err := e.db.Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (e EventRepositoryImpl) GetByID(id int) (models.Event, error) {
	var event models.Event
	err := e.db.Where("id = ?", id).First(&event).Error
	if err != nil {
		return models.Event{}, err
	}
	return event, nil
}

func (e EventRepositoryImpl) Save(event *models.Event) (models.Event, error) {
	err := e.db.Create(event).Error
	if err != nil {
		return models.Event{}, err
	}
	return *event, nil
}

func (e EventRepositoryImpl) Update(event *models.Event) (models.Event, error) {
	err := e.db.Save(event).Error
	if err != nil {
		return models.Event{}, err
	}
	return *event, nil
}

func (e EventRepositoryImpl) Delete(id int) error {
	err := e.db.Delete(&models.Event{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (e EventRepositoryImpl) GetFeaturedEvents() ([]models.Event, error) {
	var events []models.Event
	err := e.db.Where("is_featured = ?", true).Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (e EventRepositoryImpl) GetMyEvents(userId int) ([]models.Event, error) {
	// Return all events with userID in table event_users as given
	var events []models.Event
	err := e.db.Table("events").Select("events.*").Joins("join event_users on events.id = event_users.event_id").Where("event_users.user_id = ?", userId).Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil

}

func (e EventRepositoryImpl) GetCreatedEvents(userId int) ([]models.Event, error) {
	var events []models.Event
	err := e.db.Where("created_by = ?", userId).Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (e EventRepositoryImpl) BookEvent(eventID int, userID int) error {
	eventUser := models.EventUser{
		EventID: eventID,
		UserID:  userID,
	}
	err := e.db.Create(&eventUser).Error
	if err != nil {
		return err
	}
	return nil
}

func (e EventRepositoryImpl) GetBooking(eventID int, userID int) (models.EventUser, error) {
	var eventUser models.EventUser
	err := e.db.Where("event_id = ? AND user_id = ?", eventID, userID).First(&eventUser).Error
	if err != nil {
		return models.EventUser{}, err
	}
	return eventUser, nil
}

func NewEventRepository(db *gorm.DB) (*EventRepositoryImpl, error) {
	err := db.AutoMigrate(&models.Event{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&models.EventUser{})
	if err != nil {
		return nil, err
	}
	return &EventRepositoryImpl{
		db: db,
	}, nil
}
