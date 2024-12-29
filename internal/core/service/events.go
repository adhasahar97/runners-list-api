package service

import (
	"errors"
	"time"

	"github.com/aniqaqill/runners-list/internal/core/domain"
	"github.com/aniqaqill/runners-list/internal/port"
)

var (
	ErrEventDateInPast    = errors.New("event date must be in the future")
	ErrEventNameNotUnique = errors.New("event name must be unique")
)

type EventService struct {
	repo port.EventRepository
}

func NewEventService(repo port.EventRepository) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) CreateEvent(event *domain.Events) error {
	// Validate event date
	if !isEventDateInFuture(event.Date) {
		return ErrEventDateInPast
	}

	// Validate event name uniqueness
	if s.repo.EventNameExists(event.Name) {
		return ErrEventNameNotUnique
	}

	// Save the event
	return s.repo.Create(event)
}

func (s *EventService) ListEvents() ([]domain.Events, error) {
	return s.repo.FindAll()
}

func (s *EventService) DeleteEvent(id uint) error {
	event, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(event)
}

// isEventDateInFuture checks if the event date is in the future
func isEventDateInFuture(date string) bool {
	const layout = "2006-01-02"
	eventDate, err := time.Parse(layout, date)
	if err != nil {
		return false
	}
	return eventDate.After(time.Now())
}

/* The service layer contains the business logic and use cases.
It uses the interfaces (ports) defined in the port layer to interact with external systems. */
