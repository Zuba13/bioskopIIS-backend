package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "bioskop.com/projekat/bioskopIIS-backend/model"
	service "bioskop.com/projekat/bioskopIIS-backend/service"
	"github.com/gorilla/mux"
)

type ContractItemHandler struct {
	ContractItemService service.ContractItemService
}

func NewContractItemHandler(contractService service.ContractItemService) *ContractItemHandler {
	return &ContractItemHandler{
		ContractItemService: contractService,
	}
}

func (mh *ContractItemHandler) RegisterContractItemHandler(r *mux.Router) {
	r.HandleFunc("/contractItems", mh.CreateContractItem).Methods("POST")
	r.HandleFunc("/contractItems", mh.GetAllContractItems).Methods("GET")
	r.HandleFunc("/contractItems/{id}", mh.GetContractItemByID).Methods("GET")
	r.HandleFunc("/contractItems/{id}", mh.UpdateContractItem).Methods("PUT")
	r.HandleFunc("/contractItems/{id}", mh.DeleteContractItem).Methods("DELETE")
}

func (mh *ContractItemHandler) CreateContractItem(w http.ResponseWriter, r *http.Request) {
	var newContractItem model.ContractItem
	err := json.NewDecoder(r.Body).Decode(&newContractItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	contract, err := mh.ContractItemService.CreateContractItem(newContractItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusCreated, contract)
}



func (mh *ContractItemHandler) GetAllContractItems(w http.ResponseWriter, r *http.Request) {
	contractItems, err := mh.ContractItemService.GetAllContractItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, contractItems)
}

func (mh *ContractItemHandler) GetContractItemByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	contractID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	contract, err := mh.ContractItemService.GetContractItemById(uint(contractID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, contract)
}

func (mh *ContractItemHandler) UpdateContractItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	contractID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedContractItem model.ContractItem
	err = json.NewDecoder(r.Body).Decode(&updatedContractItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedContractItem.Id = uint(contractID)

	err = mh.ContractItemService.UpdateContractItem(&updatedContractItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, updatedContractItem)
}

func (mh *ContractItemHandler) DeleteContractItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	contractID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = mh.ContractItemService.DeleteContractItem(uint(contractID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "ContractItem deleted successfully"})
}
