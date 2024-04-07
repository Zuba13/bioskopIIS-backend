package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"bioskop.com/projekat/bioskopIIS-backend/model"
	"bioskop.com/projekat/bioskopIIS-backend/service"
	"github.com/gorilla/mux"
)

type TicketHandler struct {
	TicketService service.TicketService
}

func NewTicketHandler(ticketService service.TicketService) *TicketHandler {
	return &TicketHandler{
		TicketService: ticketService,
	}
}

func (th *TicketHandler) RegisterTicketHandler(r *mux.Router) {
	r.HandleFunc("/tickets", th.CreateTicket).Methods("POST")
	r.HandleFunc("/tickets", th.GetAllTickets).Methods("GET")
	r.HandleFunc("/tickets/{id}", th.GetTicketByID).Methods("GET")
	r.HandleFunc("/tickets/{id}", th.UpdateTicket).Methods("PUT")
	r.HandleFunc("/tickets/{id}", th.DeleteTicket).Methods("DELETE")
	r.HandleFunc("/tickets/user/{userID}", th.GetTicketsByUserID).Methods("GET")
}

func (th *TicketHandler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	var newTicket model.Ticket
	err := json.NewDecoder(r.Body).Decode(&newTicket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set current time for CreatedAt and UpdatedAt fields
	newTicket.CreatedAt = time.Now()
	newTicket.UpdatedAt = time.Now()

	ticket, err := th.TicketService.CreateTicket(newTicket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusCreated, ticket)
}

func (th *TicketHandler) GetAllTickets(w http.ResponseWriter, r *http.Request) {
	tickets, err := th.TicketService.GetAllTickets()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, tickets)
}

func (th *TicketHandler) GetTicketByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ticketID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ticket, err := th.TicketService.GetTicketByID(uint(ticketID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, ticket)
}

func (th *TicketHandler) UpdateTicket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ticketID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedTicket model.Ticket
	err = json.NewDecoder(r.Body).Decode(&updatedTicket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedTicket.ID = uint(ticketID)
	updatedTicket.UpdatedAt = time.Now()

	err = th.TicketService.UpdateTicket(&updatedTicket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, updatedTicket)
}

func (th *TicketHandler) DeleteTicket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ticketID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = th.TicketService.DeleteTicket(uint(ticketID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Ticket deleted successfully"})
}

func (th *TicketHandler) GetTicketsByUserID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tickets, err := th.TicketService.GetTicketsByUserID(uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, tickets)
}
