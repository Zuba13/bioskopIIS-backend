package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"bioskop.com/projekat/bioskopIIS-backend/model"
	service "bioskop.com/projekat/bioskopIIS-backend/service"
	"github.com/gorilla/mux"
)

type SupplierHandler struct {
	SupplierService service.SupplierService
}

func NewSupplierHandler(supplierService service.SupplierService) *SupplierHandler {
	return &SupplierHandler{
		SupplierService: supplierService,
	}
}

func (mh *SupplierHandler) RegisterSupplierHandler(r *mux.Router) {
	r.HandleFunc("/suppliers", mh.CreateSupplier).Methods("POST")
	r.HandleFunc("/suppliers", mh.GetAllSuppliers).Methods("GET")
	r.HandleFunc("/suppliers/{id}", mh.GetSupplierByID).Methods("GET")
	r.HandleFunc("/suppliers/{id}", mh.UpdateSupplier).Methods("PUT")
	r.HandleFunc("/suppliers/{id}", mh.DeleteSupplier).Methods("DELETE")
}

func (mh *SupplierHandler) CreateSupplier(w http.ResponseWriter, r *http.Request) {
	var newSupplier model.Supplier
	err := json.NewDecoder(r.Body).Decode(&newSupplier)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	supplier, err := mh.SupplierService.CreateSupplier(newSupplier)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusCreated, supplier)
}



func (mh *SupplierHandler) GetAllSuppliers(w http.ResponseWriter, r *http.Request) {
	suppliers, err := mh.SupplierService.GetAllSuppliers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, suppliers)
}

func (mh *SupplierHandler) GetSupplierByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	supplierID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	supplier, err := mh.SupplierService.GetSupplierById(uint(supplierID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, supplier)
}

func (mh *SupplierHandler) UpdateSupplier(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	supplierID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedSupplier model.Supplier
	err = json.NewDecoder(r.Body).Decode(&updatedSupplier)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedSupplier.Id = uint(supplierID)

	err = mh.SupplierService.UpdateSupplier(&updatedSupplier)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, updatedSupplier)
}

func (mh *SupplierHandler) DeleteSupplier(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	supplierID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = mh.SupplierService.DeleteSupplier(uint(supplierID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Supplier deleted successfully"})
}

