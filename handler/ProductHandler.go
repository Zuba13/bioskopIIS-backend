package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "bioskop.com/projekat/bioskopIIS-backend/model"
	service "bioskop.com/projekat/bioskopIIS-backend/service"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	ProductService service.ProductService
}

func NewProductHandler(contractService service.ProductService) *ProductHandler {
	return &ProductHandler{
		ProductService: contractService,
	}
}

func (mh *ProductHandler) RegisterProductHandler(r *mux.Router) {
	r.HandleFunc("/products", mh.CreateProduct).Methods("POST")
	r.HandleFunc("/products", mh.GetAllProducts).Methods("GET")
	r.HandleFunc("/products/{id}", mh.GetProductById).Methods("GET")
	r.HandleFunc("/products/{id}", mh.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", mh.DeleteProduct).Methods("DELETE")
}

func (mh *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct model.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	contract, err := mh.ProductService.CreateProduct(newProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusCreated, contract)
}



func (mh *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := mh.ProductService.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

func (mh *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	contractID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	contract, err := mh.ProductService.GetProductById(uint(contractID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, contract)
}

func (mh *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedProduct model.Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedProduct.Id = uint(id)

	err = mh.ProductService.UpdateProduct(&updatedProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, updatedProduct)
}

func (mh *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	contractID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = mh.ProductService.DeleteProduct(uint(contractID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Product deleted successfully"})
}
