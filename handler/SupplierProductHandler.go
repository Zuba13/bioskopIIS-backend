package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "bioskop.com/projekat/bioskopIIS-backend/model"
	service "bioskop.com/projekat/bioskopIIS-backend/service"
	"github.com/gorilla/mux"
)

type SupplierProductHandler struct {
	SupplierProductService service.SupplierProductService
}

func NewSupplierProductHandler(menuService service.SupplierProductService) *SupplierProductHandler {
	return &SupplierProductHandler{
		SupplierProductService: menuService,
	}
}

func (mh *SupplierProductHandler) RegisterSupplierProductHandler(r *mux.Router) {
	r.HandleFunc("/supplierProducts", mh.CreateSupplierProduct).Methods("POST")
	r.HandleFunc("/supplierProducts", mh.GetAllSupplierProducts).Methods("GET")
	r.HandleFunc("/supplierProducts/{id}", mh.GetSupplierProductById).Methods("GET")
	//r.HandleFunc("/supplierProducts/{id}", mh.UpdateSupplierProduct).Methods("PUT")
	r.HandleFunc("/supplierProducts/{id}", mh.DeleteSupplierProduct).Methods("DELETE")
}

func (mh *SupplierProductHandler) CreateSupplierProduct(w http.ResponseWriter, r *http.Request) {
	var newSupplierProduct model.SupplierProduct
	err := json.NewDecoder(r.Body).Decode(&newSupplierProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	stockItem, err := mh.SupplierProductService.CreateSupplierProduct(newSupplierProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusCreated, stockItem)
}



func (mh *SupplierProductHandler) GetAllSupplierProducts(w http.ResponseWriter, r *http.Request) {
	supplierProducts, err := mh.SupplierProductService.GetAllSupplierProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, supplierProducts)
}

func (mh *SupplierProductHandler) GetSupplierProductById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	menuID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	menu, err := mh.SupplierProductService.GetSupplierProductById(uint(menuID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, menu)
}

// func (mh *SupplierProductHandler) UpdateSupplierProduct(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	menuID, err := strconv.ParseUint(params["id"], 10, 64)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	var updatedSupplierProduct model.SupplierProduct
// 	err = json.NewDecoder(r.Body).Decode(&updatedSupplierProduct)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	updatedSupplierProduct.Id = uint(menuID)

// 	err = mh.SupplierProductService.UpdateSupplierProduct(&updatedSupplierProduct)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, updatedSupplierProduct)
// }

func (mh *SupplierProductHandler) DeleteSupplierProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	menuId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = mh.SupplierProductService.DeleteSupplierProduct(uint(menuId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "SupplierProduct deleted successfully"})
}
