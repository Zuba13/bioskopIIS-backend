package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

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
	r.HandleFunc("/contracts", mh.CreateContract).Methods("POST")
	r.HandleFunc("/contracts", mh.GetAllContracts).Methods("GET")
	r.HandleFunc("/suppliercontracts/{supplier_id}", mh.GetSupplierContracts).Methods("GET")
	r.HandleFunc("/contracts/{id}", mh.GetContractByID).Methods("GET")
	r.HandleFunc("/contracts/{id}", mh.UpdateContract).Methods("PUT")
	r.HandleFunc("/contracts/{id}", mh.DeleteContract).Methods("DELETE")
}

func (ch *ContractHandler) CreateContract(w http.ResponseWriter, r *http.Request) {
	var createContractDTO model.CreateContractDTO

	err := json.NewDecoder(r.Body).Decode(&createContractDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	supplierId, err := strconv.Atoi(createContractDTO.BaseData.SupplierId)
	if err != nil {
		http.Error(w, "Invalid supplier ID", http.StatusBadRequest)
		return
	}

	contractType, err := strconv.Atoi(createContractDTO.BaseData.ContractType)
	if err != nil {
		http.Error(w, "Invalid supplier ID", http.StatusBadRequest)
		return
	}
	
	newContract := model.Contract{
		Name:            createContractDTO.BaseData.Name,
		ValidFrom:       createContractDTO.BaseData.ValidFrom,
		ValidUntil:      createContractDTO.BaseData.ValidUntil,
		SupplierId:      uint(supplierId),
		DateOfSignature: time.Now(), // or however you get this value
		ContractType:    model.Type(contractType),
		ContractItems:   mapContractItems(createContractDTO.ContractItems),
	}
	
	contract, err := ch.ContractService.CreateContract(newContract)
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

func (mh *ContractHandler) GetAllWeeklyContracts(w http.ResponseWriter, r *http.Request) {
	contracts, err := mh.ContractService.GetTodayyWeeklyContracts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, contracts)
}

func (mh *ContractHandler) GetAllAtOnceContracts(w http.ResponseWriter, r *http.Request) {
	contracts, err := mh.ContractService.GetTodayAtOnceContracts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, contracts)
}

func (mh *ContractHandler) GetSupplierContracts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	supplierId, err := strconv.ParseUint(params["supplier_id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	contracts, err := mh.ContractService.GetAllSupplierContracts(uint(supplierId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
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

func mapContractItems(dtoItems []model.ContractItemDTO) []model.ContractItem {
	var items []model.ContractItem
	for _, dtoItem := range dtoItems {
		item := model.ContractItem{
			ProductId:     dtoItem.ProductId,
			Quantity: dtoItem.Quantity,
			Price:    dtoItem.Price,
		}
		items = append(items, item)
	}
	return items
}