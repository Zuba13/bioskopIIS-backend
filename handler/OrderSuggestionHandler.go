package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "bioskop.com/projekat/bioskopIIS-backend/model"
	service "bioskop.com/projekat/bioskopIIS-backend/service"
	"github.com/gorilla/mux"
)

type OrderSuggestionHandler struct {
	OrderSuggestionService service.OrderSuggestionService
}

func NewOrderSuggestionHandler(menuService service.OrderSuggestionService) *OrderSuggestionHandler {
	return &OrderSuggestionHandler{
		OrderSuggestionService: menuService,
	}
}

func (mh *OrderSuggestionHandler) RegisterOrderSuggestionHandler(r *mux.Router) {
	r.HandleFunc("/orderSuggestions", mh.CreateOrderSuggestion).Methods("POST")
	r.HandleFunc("/orderSuggestions", mh.GetAllOrderSuggestions).Methods("GET")
	r.HandleFunc("/orderSuggestions/{id}", mh.GetOrderSuggestionById).Methods("GET")
	r.HandleFunc("/orderSuggestions/{id}", mh.UpdateOrderSuggestion).Methods("PUT")
	r.HandleFunc("/orderSuggestions/{id}", mh.DeleteOrderSuggestion).Methods("DELETE")
}

func (mh *OrderSuggestionHandler) CreateOrderSuggestion(w http.ResponseWriter, r *http.Request) {
	var newOrderSuggestion model.OrderSuggestion
	err := json.NewDecoder(r.Body).Decode(&newOrderSuggestion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	menu, err := mh.OrderSuggestionService.CreateOrderSuggestion(newOrderSuggestion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusCreated, menu)
}



func (mh *OrderSuggestionHandler) GetAllOrderSuggestions(w http.ResponseWriter, r *http.Request) {
	orderSuggestions, err := mh.OrderSuggestionService.GetAllOrderSuggestions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, orderSuggestions)
}

func (mh *OrderSuggestionHandler) GetOrderSuggestionById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	menuID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	menu, err := mh.OrderSuggestionService.GetOrderSuggestionById(uint(menuID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, menu)
}

func (mh *OrderSuggestionHandler) UpdateOrderSuggestion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	menuID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedOrderSuggestion model.OrderSuggestion
	err = json.NewDecoder(r.Body).Decode(&updatedOrderSuggestion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedOrderSuggestion.Id = uint(menuID)

	err = mh.OrderSuggestionService.UpdateOrderSuggestion(&updatedOrderSuggestion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, updatedOrderSuggestion)
}

func (mh *OrderSuggestionHandler) DeleteOrderSuggestion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	menuId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = mh.OrderSuggestionService.DeleteOrderSuggestion(uint(menuId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "OrderSuggestion deleted successfully"})
}
