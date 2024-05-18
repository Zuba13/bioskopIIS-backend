package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"bioskop.com/projekat/bioskopIIS-backend/model"
	"bioskop.com/projekat/bioskopIIS-backend/service"
	"github.com/dgrijalva/jwt-go"
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
	r.HandleFunc("/distribution/contract", dch.CreateContract).Methods("POST")
	r.HandleFunc("/distribution/contract", dch.UpdateContract).Methods("PUT")
	r.HandleFunc("/distribution/company", dch.GetAllCompanies).Methods("GET")
}

func (dch *DistributionContractHandler) CreateContract(w http.ResponseWriter, r *http.Request) {
	var contract model.DistributionContract
	if err := json.NewDecoder(r.Body).Decode(&contract); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userId, err := extractUserIDFromJWT(r)
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

func extractUserIDFromJWT(r *http.Request) (uint, error) {
	// Retrieve the token from the Authorization header
	authHeader := r.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JWTSecretKey, nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIdFloat, ok := claims["user_id"].(float64)
		if !ok {
			return 0, fmt.Errorf("user_id is not a number")
		}
		userId := uint(userIdFloat)
		return userId, nil
	} else {
		return 0, fmt.Errorf("invalid token")
	}
}
