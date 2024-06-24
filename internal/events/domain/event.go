package domain

import (
	"errors"
	"time"
)

var (
	ErrInvalidEvent    = errors.New("invalid event data")
	ErrEventFull       = errors.New("event is full")
	ErrTicketNotFound  = errors.New("ticket not found")
	ErrTicketNotEnough = errors.New("not enough tickets available")
	ErrEventNotFound   = errors.New("event not found")
)

type Rating string

const (
	RatingLivre Rating = "L"
	Rating10    Rating = "L10"
	Rating12    Rating = "L12"
	Rating14    Rating = "L14"
	Rating16    Rating = "L16"
	Rating18    Rating = "L18"
)

type Event struct {
	ID           string
	Name         string
	Location     string
	Organization string
	Rating       Rating
	Date         time.Time
	ImageURL     string
	Capacity     int
	Price        float64
	PartnerID    int
	Spots        []Spot
	Tickets      []Ticket
}

// Validate checks if the event data is valid.
func (e *Event) Validate() error {
	if e.Name == "" {
		return errors.New("event name is required")
	}
	if e.Date.Before(time.Now()) {
		return errors.New("event date must be in the future")
	}
	if e.Capacity <= 0 {
		return errors.New("event capacity must be greater than zero")
	}
	if e.Price <= 0 {
		return errors.New("event price must be greater than zero")
	}

	return nil
}

// AddSpot adds a spot to the event.
func (e *Event) AddSpot(name string) (*Spot, error) {
	spot, err := NewSpot(e, name)
	if err != nil {
		return nil, err
	}
	e.Spots = append(e.Spots, *spot)
	return spot, nil
}
