package domain

import "errors"

type TicketKind string

var ErrTicketPriceZero = errors.New("ticket price must be greater than zero")

const (
	TicketKindHalf TicketKind = "half"
	TicketKindFull TicketKind = "full"
)

type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot
	TicketKind TicketKind
	Price      float64
}

func IsValidTicketType(ticketKind TicketKind) bool {
	return ticketKind == TicketKindHalf || ticketKind == TicketKindFull
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
