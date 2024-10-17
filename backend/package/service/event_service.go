package service

import (
	"fmt"

	"github.com/HermanPlay/web-app-backend/package/domain/models"
	"github.com/HermanPlay/web-app-backend/package/domain/schemas"
	"github.com/HermanPlay/web-app-backend/package/repository"
	"gorm.io/gorm"
)

type EventService interface {
	GetAllEvent() ([]*schemas.Event, error)
	GetEventByID(id int) (*schemas.Event, error)
	CreateEvent(event *schemas.EventInput, createdBy int) (*schemas.Event, error)
	UpdateEvent(event *schemas.EventUpdate, id int) (*schemas.Event, error)
	DeleteEvent(id int) error
	GetFeaturedEvents() ([]*schemas.Event, error)
	GetMyEvents(userId int) ([]*schemas.Event, error)
	BookEvent(eventID int, userID int) error
}

var (
	ErrBookingExists = fmt.Errorf("booking already exists")
)

type EventServiceImpl struct {
	eventRepository repository.EventRepository
}

func (e EventServiceImpl) GetAllEvent() ([]*schemas.Event, error) {
	events, err := e.eventRepository.GetAll()
	if err != nil {
		return nil, err
	}
	eventResponse := []*schemas.Event{}
	for _, event := range events {
		eventResponse = append(eventResponse, e.createEventResponse(&event))
	}

	return eventResponse, nil
}

func (e EventServiceImpl) GetEventByID(id int) (*schemas.Event, error) {
	event, err := e.eventRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	eventResponse := e.createEventResponse(&event)
	return eventResponse, nil
}

func (e EventServiceImpl) CreateEvent(eventInput *schemas.EventInput, createdBy int) (*schemas.Event, error) {
	if len(eventInput.ShortDescription) > 100 {
		return nil, ErrInputTooLong
	}
	modelEvent := e.createEventModel(eventInput)
	modelEvent.CreatedBy = createdBy
	event, err := e.eventRepository.Save(modelEvent)
	if err != nil {
		return nil, err
	}
	eventResponse := e.createEventResponse(&event)

	return eventResponse, nil

}

func (e EventServiceImpl) UpdateEvent(eventUpdate *schemas.EventUpdate, id int) (*schemas.Event, error) {
	eventModel, err := e.eventRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	e.updateModel(&eventModel, eventUpdate)

	event, err := e.eventRepository.Update(&eventModel)

	if err != nil {
		return nil, err
	}

	eventResponse := e.createEventResponse(&event)

	return eventResponse, nil
}

func (e EventServiceImpl) DeleteEvent(id int) error {
	return e.eventRepository.Delete(id)
}

func (e EventServiceImpl) GetFeaturedEvents() ([]*schemas.Event, error) {
	events, err := e.eventRepository.GetFeaturedEvents()
	if err != nil {
		return nil, err
	}
	eventResponse := []*schemas.Event{}
	for _, event := range events {
		eventResponse = append(eventResponse, e.createEventResponse(&event))
	}

	return eventResponse, nil
}

func (e EventServiceImpl) GetMyEvents(userId int) ([]*schemas.Event, error) {
	events, err := e.eventRepository.GetMyEvents(userId)
	if err != nil {
		return nil, err
	}
	createdEvents, err := e.eventRepository.GetCreatedEvents(userId)
	if err != nil {
		return nil, err
	}
	events = append(events, createdEvents...)
	eventResponse := []*schemas.Event{}
	for _, event := range events {
		eventResponse = append(eventResponse, e.createEventResponse(&event))
	}
	return eventResponse, nil

}

func (e EventServiceImpl) BookEvent(eventID int, userID int) error {
	_, err := e.eventRepository.GetBooking(eventID, userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return e.eventRepository.BookEvent(eventID, userID)
		}
		return err
	}
	return ErrBookingExists

}

func (e EventServiceImpl) createEventModel(event *schemas.EventInput) *models.Event {
	return &models.Event{
		Title:            event.Title,
		ShortDescription: event.ShortDescription,
		Description:      event.Description,
		Location:         event.Location,
		Date:             event.Date,
		Time:             event.Time,
		IsFeatured:       event.IsFeatured,
	}
}

func (e EventServiceImpl) updateModel(eventModel *models.Event, eventUpdate *schemas.EventUpdate) {
	if eventUpdate.Title != "" {
		eventModel.Title = eventUpdate.Title
	}
	if eventUpdate.ShortDescription != "" {
		eventModel.ShortDescription = eventUpdate.ShortDescription
	}
	if eventUpdate.Description != "" {
		eventModel.Description = eventUpdate.Description
	}
	if eventUpdate.Location != "" {
		eventModel.Location = eventUpdate.Location
	}
	if eventUpdate.Date != "" {
		eventModel.Date = eventUpdate.Date
	}
	if eventUpdate.Time != "" {
		eventModel.Time = eventUpdate.Time
	}
	if eventUpdate.IsFeatured != eventModel.IsFeatured {
		eventModel.IsFeatured = eventUpdate.IsFeatured
	}

}

func (e EventServiceImpl) createEventResponse(event *models.Event) *schemas.Event {
	return &schemas.Event{
		ID:               event.ID,
		Title:            event.Title,
		ShortDescription: event.ShortDescription,
		Description:      event.Description,
		Location:         event.Location,
		Date:             event.Date,
		Time:             event.Time,
		IsFeatured:       event.IsFeatured,
		CreatedBy:        event.CreatedBy,
	}
}
func NewEventService(eventRepository repository.EventRepository) EventService {
	return &EventServiceImpl{
		eventRepository: eventRepository,
	}
}
