package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	model "bioskop.com/projekat/bioskopIIS-backend/model"
	service "bioskop.com/projekat/bioskopIIS-backend/service"
	"github.com/gorilla/mux"
)

type ContractHandler struct {
	ContractService service.ContractService
	ContractItemService service.ContractItemService
}

func NewContractHandler(contractService service.ContractService, contractItemService service.ContractItemService) *ContractHandler {
	return &ContractHandler{
		ContractService: contractService,
		ContractItemService: contractItemService,
	}
}

func (mh *ContractHandler) RegisterContractHandler(r *mux.Router) {
	r.HandleFunc("/contracts", mh.CreateContractWithItems).Methods("POST")
	//r.HandleFunc("/contracts", mh.CreateContract).Methods("POST")
	r.HandleFunc("/contracts", mh.GetAllContracts).Methods("GET")
	r.HandleFunc("/contracts/{id}", mh.GetContractByID).Methods("GET")
	r.HandleFunc("/contracts/{id}", mh.UpdateContract).Methods("PUT")
	r.HandleFunc("/contracts/{id}", mh.DeleteContract).Methods("DELETE")
}

func (ch *ContractHandler) CreateContract(w http.ResponseWriter, r *http.Request) {
	var newContract model.Contract
	err := json.NewDecoder(r.Body).Decode(&newContract)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	contract, err := ch.ContractService.CreateContract(newContract)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusCreated, contract)
} 

func (ch *ContractHandler) CreateContractWithItems(w http.ResponseWriter, r *http.Request) {
	var newContract model.Contract
	err := json.NewDecoder(r.Body).Decode(&newContract)
	fmt.Println(r.Body)
	fmt.Println(newContract)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var conItems []model.ContractItem
	err = json.NewDecoder(r.Body).Decode(&conItems)
	fmt.Println(&conItems)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	contract, err := ch.ContractService.CreateContract(newContract)
	if err != nil {
		fmt.Println("Puklo")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, item := range conItems {
		// Do something with each contract item, e.g., validate or process
		// item is of type model.ContractItem
		contractItem, err := ch.ContractItemService.CreateContractItem(item)
		fmt.Println("Usao u iteme")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Append the newly created contract item to the contract object
		contract.ContractItems = append(contract.ContractItems, contractItem)
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
