package http

import (
	"encoding/json"
	"net/http"

	"github.com/laismatie/golang-events-api/internal/events/usecase"
)

// EventsHandler handles HTTP requests for events.
type EventsHandler struct {
	listEventsUseCase *usecase.ListEventsUseCase
	listSpotsUseCase  *usecase.ListSpotsUseCase
	getEventUseCase   *usecase.GetEventUseCase
	buyTicketsUseCase *usecase.BuyTicketsUseCase
}

// NewEventsHandler creates a new EventsHandler.
func NewEventsHandler(
	listEventsUseCase *usecase.ListEventsUseCase,
	listSpotsUseCase *usecase.ListSpotsUseCase,
	getEventUseCase *usecase.GetEventUseCase,
	buyTicketsUseCase *usecase.BuyTicketsUseCase,
) *EventsHandler {
	return &EventsHandler{
		listEventsUseCase: listEventsUseCase,
		listSpotsUseCase:  listSpotsUseCase,
		getEventUseCase:   getEventUseCase,
		buyTicketsUseCase: buyTicketsUseCase,
	}
}

// ListEvents handles the request to list all events.
func (h *EventsHandler) ListEvents(w http.ResponseWriter, r *http.Request) {
	output, err := h.listEventsUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

// GetEvent handles the request to get details of a specific event.
func (h *EventsHandler) GetEvent(w http.ResponseWriter, r *http.Request) {
	eventID := r.PathValue("eventID")
	input := usecase.GetEventInputDTO{ID: eventID}

	output, err := h.getEventUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

// ListSpots lists spots for an event.
func (h *EventsHandler) ListSpots(w http.ResponseWriter, r *http.Request) {
	eventID := r.PathValue("eventID")
	input := usecase.ListSpotsInputDTO{EventID: eventID}

	output, err := h.listSpotsUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

// BuyTickets handles the request to buy tickets for an event.
func (h *EventsHandler) BuyTickets(w http.ResponseWriter, r *http.Request) {
	var input usecase.BuyTicketsInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.buyTicketsUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
