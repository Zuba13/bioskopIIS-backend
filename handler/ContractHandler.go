package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "bioskop.com/projekat/bioskopIIS-backend/model"
	service "bioskop.com/projekat/bioskopIIS-backend/service"
	"github.com/gorilla/mux"
)

type ContractHandler struct {
	ContractService service.ContractService
}

func NewContractHandler(contractService service.ContractService) *ContractHandler {
	return &ContractHandler{
		ContractService: contractService,
	}
}

func (mh *ContractHandler) RegisterContractHandler(r *mux.Router) {
	r.HandleFunc("/contracts", mh.CreateContract).Methods("POST")
	r.HandleFunc("/contracts", mh.GetAllContracts).Methods("GET")
	r.HandleFunc("/contracts/{id}", mh.GetContractByID).Methods("GET")
	r.HandleFunc("/contracts/{id}", mh.UpdateContract).Methods("PUT")
	r.HandleFunc("/contracts/{id}", mh.DeleteContract).Methods("DELETE")
}

func (mh *ContractHandler) CreateContract(w http.ResponseWriter, r *http.Request) {
	var newContract model.Contract
	err := json.NewDecoder(r.Body).Decode(&newContract)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	contract, err := mh.ContractService.CreateContract(newContract)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusCreated, contract)
}



func (mh *ContractHandler) GetAllContracts(w http.ResponseWriter, r *http.Request) {
	contracts, err := mh.ContractService.GetAllContracts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, contracts)
}

func (mh *ContractHandler) GetContractByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	contractID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	contract, err := mh.ContractService.GetContractById(uint(contractID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, contract)
}

func (mh *ContractHandler) UpdateContract(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	contractID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedContract model.Contract
	err = json.NewDecoder(r.Body).Decode(&updatedContract)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedContract.Id = uint(contractID)

	err = mh.ContractService.UpdateContract(&updatedContract)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, updatedContract)
}

func (mh *ContractHandler) DeleteContract(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	contractID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = mh.ContractService.DeleteContract(uint(contractID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Contract deleted successfully"})
}
