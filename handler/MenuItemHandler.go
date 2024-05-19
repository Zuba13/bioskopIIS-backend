package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "bioskop.com/projekat/bioskopIIS-backend/model"
	service "bioskop.com/projekat/bioskopIIS-backend/service"
	"github.com/gorilla/mux"
)

type MenuItemHandler struct {
	MenuItemService service.MenuItemService
}

func NewMenuItemHandler(menuService service.MenuItemService) *MenuItemHandler {
	return &MenuItemHandler{
		MenuItemService: menuService,
	}
}

func (mh *MenuItemHandler) RegisterMenuItemHandler(r *mux.Router) {
	r.HandleFunc("/menuItems", mh.CreateMenuItem).Methods("POST")
	r.HandleFunc("/menuItems", mh.GetAllMenuItems).Methods("GET")
	r.HandleFunc("/menuItems/{id}", mh.GetMenuItemById).Methods("GET")
	r.HandleFunc("/menuItems/{id}", mh.UpdateMenuItem).Methods("PUT")
	r.HandleFunc("/menuItems/{id}", mh.DeleteMenuItem).Methods("DELETE")
}

func (mh *MenuItemHandler) CreateMenuItem(w http.ResponseWriter, r *http.Request) {
	var newMenuItem model.MenuItem
	err := json.NewDecoder(r.Body).Decode(&newMenuItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	menu, err := mh.MenuItemService.CreateMenuItem(newMenuItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusCreated, menu)
}



func (mh *MenuItemHandler) GetAllMenuItems(w http.ResponseWriter, r *http.Request) {
	menuItems, err := mh.MenuItemService.GetAllMenuItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, menuItems)
}

func (mh *MenuItemHandler) GetMenuItemById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	menuID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	menu, err := mh.MenuItemService.GetMenuItemById(uint(menuID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, menu)
}

func (mh *MenuItemHandler) UpdateMenuItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	menuID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedMenuItem model.MenuItem
	err = json.NewDecoder(r.Body).Decode(&updatedMenuItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedMenuItem.Id = uint(menuID)

	err = mh.MenuItemService.UpdateMenuItem(&updatedMenuItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, updatedMenuItem)
}

func (mh *MenuItemHandler) DeleteMenuItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	menuId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = mh.MenuItemService.DeleteMenuItem(uint(menuId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "MenuItem deleted successfully"})
}
