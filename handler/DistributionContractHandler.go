package handler

import (
	"encoding/json"
	"net/http"

	"bioskop.com/projekat/bioskopIIS-backend/model"
	"bioskop.com/projekat/bioskopIIS-backend/service"
	"github.com/gorilla/mux"
)

type DistributionContractHandler struct {
	DistributionContractService *service.DistributionContractService
}

func NewDistributionContractHandler(distributionContractService *service.DistributionContractService) *DistributionContractHandler {
	return &DistributionContractHandler{
		DistributionContractService: distributionContractService,
	}
}

func (dch *DistributionContractHandler) RegisterDistributionContractHandler(r *mux.Router) {
	r.HandleFunc("/distribution/contract", RequireRole("manager", dch.CreateContract)).Methods("POST")
	r.HandleFunc("/distribution/contract", RequireRole("manager", dch.UpdateContract)).Methods("PUT")
	r.HandleFunc("/distribution/company", RequireRole("manager", dch.GetAllCompanies)).Methods("GET")
}

func (dch *DistributionContractHandler) CreateContract(w http.ResponseWriter, r *http.Request) {
	var contract model.DistributionContract
	if err := json.NewDecoder(r.Body).Decode(&contract); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userId, err := ExtractUserIDFromJWT(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	createdContract, err := dch.DistributionContractService.CreateContract(&contract, userId)
	if err != nil {
		customErr, ok := err.(*service.CustomError)
		if ok {
			http.Error(w, customErr.Message, service.ErrorCodeToHTTPStatus(customErr.Code))
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	respondWithJSON(w, http.StatusCreated, createdContract)
}

func (dch *DistributionContractHandler) UpdateContract(w http.ResponseWriter, r *http.Request) {
	var contract model.DistributionContract
	if err := json.NewDecoder(r.Body).Decode(&contract); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedContract, err := dch.DistributionContractService.UpdateContract(&contract)
	if err != nil {
		customErr, ok := err.(*service.CustomError)
		if ok {
			http.Error(w, customErr.Message, service.ErrorCodeToHTTPStatus(customErr.Code))
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	respondWithJSON(w, http.StatusOK, updatedContract)
}

func (dch *DistributionContractHandler) GetAllCompanies(w http.ResponseWriter, r *http.Request) {
	companies, err := dch.DistributionContractService.GetAllCompanies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, companies)
}
