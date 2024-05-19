package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "bioskop.com/projekat/bioskopIIS-backend/model"
	service "bioskop.com/projekat/bioskopIIS-backend/service"
	"github.com/gorilla/mux"
)

type StockItemHandler struct {
	StockItemService service.StockItemService
}

func NewStockItemHandler(menuService service.StockItemService) *StockItemHandler {
	return &StockItemHandler{
		StockItemService: menuService,
	}
}

func (mh *StockItemHandler) RegisterStockItemHandler(r *mux.Router) {
	r.HandleFunc("/stockItems", mh.CreateStockItem).Methods("POST")
	r.HandleFunc("/stockItems", mh.GetAllStockItems).Methods("GET")
	r.HandleFunc("/stockItems/{id}", mh.GetStockItemById).Methods("GET")
	r.HandleFunc("/stockItems/{id}", mh.UpdateStockItem).Methods("PUT")
	r.HandleFunc("/stockItems/{id}", mh.DeleteStockItem).Methods("DELETE")
}

func (mh *StockItemHandler) CreateStockItem(w http.ResponseWriter, r *http.Request) {
	var newStockItem model.StockItem
	err := json.NewDecoder(r.Body).Decode(&newStockItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	stockItem, err := mh.StockItemService.CreateStockItem(newStockItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusCreated, stockItem)
}



func (mh *StockItemHandler) GetAllStockItems(w http.ResponseWriter, r *http.Request) {
	stockItems, err := mh.StockItemService.GetAllStockItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, stockItems)
}

func (mh *StockItemHandler) GetStockItemById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	menuID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	menu, err := mh.StockItemService.GetStockItemById(uint(menuID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, menu)
}

func (mh *StockItemHandler) UpdateStockItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	menuID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedStockItem model.StockItem
	err = json.NewDecoder(r.Body).Decode(&updatedStockItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedStockItem.Id = uint(menuID)

	err = mh.StockItemService.UpdateStockItem(&updatedStockItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, updatedStockItem)
}

func (mh *StockItemHandler) DeleteStockItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	menuId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = mh.StockItemService.DeleteStockItem(uint(menuId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "StockItem deleted successfully"})
}
