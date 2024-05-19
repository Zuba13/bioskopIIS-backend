package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "bioskop.com/projekat/bioskopIIS-backend/model"
	service "bioskop.com/projekat/bioskopIIS-backend/service"
	"github.com/gorilla/mux"
)

type MenuHandler struct {
	MenuService service.MenuService
	MenuItemService service.MenuItemService
}

func NewMenuHandler(menuService service.MenuService, menuItemService service.MenuItemService) *MenuHandler {
	return &MenuHandler{
		MenuService: menuService,
		MenuItemService: menuItemService,
	}
}

func (mh *MenuHandler) RegisterMenuHandler(r *mux.Router) {
	r.HandleFunc("/menus", mh.CreateMenu).Methods("POST")
	r.HandleFunc("/menus", mh.GetAllMenus).Methods("GET")
	r.HandleFunc("/menus/{id}", mh.GetMenuById).Methods("GET")
	r.HandleFunc("/menus/{id}", mh.UpdateMenu).Methods("PUT")
	r.HandleFunc("/menus/{id}", mh.DeleteMenu).Methods("DELETE")
}

func (ch *MenuHandler) CreateMenu(w http.ResponseWriter, r *http.Request) {
	var createMenu model.Menu

	err := json.NewDecoder(r.Body).Decode(&createMenu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	
	menu, err := ch.MenuService.CreateMenu(createMenu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusCreated, menu)
} 


func (mh *MenuHandler) GetAllMenus(w http.ResponseWriter, r *http.Request) {
	menus, err := mh.MenuService.GetAllMenus()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, menus)
}



func (mh *MenuHandler) GetMenuById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	menuID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	menu, err := mh.MenuService.GetMenuById(uint(menuID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, menu)
}

func (mh *MenuHandler) UpdateMenu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	menuID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedMenu model.Menu
	err = json.NewDecoder(r.Body).Decode(&updatedMenu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedMenu.Id = uint(menuID)

	err = mh.MenuService.UpdateMenu(&updatedMenu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, updatedMenu)
}

func (mh *MenuHandler) DeleteMenu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	menuID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = mh.MenuService.DeleteMenu(uint(menuID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Menu deleted successfully"})
}

