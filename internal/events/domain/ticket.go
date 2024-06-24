package domain

import "errors"

var (
	ErrInvalidTicketKind = errors.New("invalid ticket kind")
	ErrTicketPriceZero   = errors.New("ticket price must be greater than zero")
)

type TicketKind string

const (
	TicketKindHalf TicketKind = "half"
	TicketKindFull TicketKind = "full"
)

func IsValidTicketType(ticketKind TicketKind) bool {
	return ticketKind == TicketKindHalf || ticketKind == TicketKindFull
}

type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot
	TicketKind TicketKind
	Price      float64
}

// CalculatePrice calculates the price based on the ticket kind.
func (t *Ticket) CalculatePrice() {
	if t.TicketKind == TicketKindHalf {
		t.Price /= 2
	}
}

func (t *Ticket) Validate() error {
	if t.Price <= 0 {
		return ErrTicketPriceZero
	}

	return nil
}
