package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"bioskop.com/projekat/bioskopIIS-backend/model"
	"bioskop.com/projekat/bioskopIIS-backend/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
    UserService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
    return &UserHandler{
			UserService: userService,
    }
}

func (uh *UserHandler) RegisterUserHandler(r *mux.Router) {
    r.HandleFunc("/users/register", uh.RegisterUser).Methods("POST")
		r.HandleFunc("/users/login", uh.LoginUser).Methods("POST")
    r.HandleFunc("/users/{id}", uh.GetUserByID).Methods("GET")
    r.HandleFunc("/users/{id}", uh.UpdateUser).Methods("PUT")
    r.HandleFunc("/users/{id}", uh.DeleteUser).Methods("DELETE")
}

var JWTSecretKey = []byte("kljuc_za_jwt")

type LoginResponse struct {
    Token string `json:"token"`
}

type LoginRequestBody struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func (uh *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var newUser model.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
	}

	// Check if all required fields are provided
	if newUser.Username == "" || newUser.Password == "" || newUser.Email == "" || newUser.FirstName == "" || newUser.LastName == "" {
			fmt.Println(newUser)
			http.Error(w, "All fields are required!", http.StatusBadRequest)
			return
	}

	// Perform user registration
	user, err := uh.UserService.RegisterUser(newUser.Username, newUser.Email, newUser.Password, newUser.FirstName, newUser.LastName)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	respondWithJSON(w, http.StatusCreated, user)
}


func (uh *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var requestBody LoginRequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
	}

	user, err := uh.UserService.GetUserByUsername(requestBody.Username)
	if err != nil {
			http.Error(w, "Invalid username or password 1", http.StatusUnauthorized)
			return
	}

	if !verifyPassword(user.Password, requestBody.Password) {
		fmt.Println(user.Password)
		fmt.Println(requestBody.Password)
			http.Error(w, "Invalid username or password 2", http.StatusUnauthorized)
			return
	}

	token, err := generateJWTToken(user.ID)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	response := LoginResponse{Token: token}
	respondWithJSON(w, http.StatusOK, response)
}



func (uh *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    userID, err := strconv.ParseUint(params["id"], 10, 64)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    user, err := uh.UserService.GetUserByID(uint(userID))
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    respondWithJSON(w, http.StatusOK, user)
}

func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    userID, err := strconv.ParseUint(params["id"], 10, 64)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var updatedUser model.User
    err = json.NewDecoder(r.Body).Decode(&updatedUser)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    updatedUser.ID = uint(userID)

    err = uh.UserService.UpdateUser(&updatedUser)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    respondWithJSON(w, http.StatusOK, updatedUser)
}

func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
		if !isAllowedRole(r, "admin", uh.UserService) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

    params := mux.Vars(r)
    userID, err := strconv.ParseUint(params["id"], 10, 64)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = uh.UserService.DeleteUser(uint(userID))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    respondWithJSON(w, http.StatusOK, map[string]string{"message": "User deleted successfully"})
}

func verifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func generateJWTToken(userID uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() 

	tokenString, err := token.SignedString(JWTSecretKey)
	if err != nil {
			return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*Claims, error) {
	// Parse the JWT token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return JWTSecretKey, nil // JWTSecretKey should be your secret key used for signing tokens
	})
	if err != nil {
			return nil, err
	}

	// Check if token is valid
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
			return nil, errors.New("invalid token")
	}

	return claims, nil
}

// RequiresRole wraps a handler function with role-based access control
func RequiresRole(handler http.HandlerFunc, requiredRole string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
			// Extract JWT token from Authorization header
			tokenString := extractTokenFromHeader(r)
			if tokenString == "" {
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
			}

			// Parse JWT token
			claims, err := ParseToken(tokenString)
			if err != nil {
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
			}

			// Check if user has required role
			if claims.Role != requiredRole {
					http.Error(w, "Forbidden", http.StatusForbidden)
					return
			}

			handler(w, r)
	}
}

// extractTokenFromHeader extracts the JWT token from the Authorization header
func extractTokenFromHeader(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
			return ""
	}

	// Check if Authorization header is in the format "Bearer <token>"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
			return ""
	}

	return parts[1]
}

func isAllowedRole(r *http.Request, requiredRole string, userService service.UserService) bool {
	// Extract JWT token from Authorization header
	tokenString := extractTokenFromHeader(r)
	if tokenString == "" {
			return false
	}

	// Parse JWT token to get the user ID
	claims, err := ParseToken(tokenString)
	if err != nil {
			return false
	}


	// Get the user from the database based on the user ID
	user, err := userService.GetUserByID(claims.UserID)
	if err != nil {
			return false
	}

	userRole := strings.ReplaceAll(user.Role, " ", "")
  userRole = strings.ReplaceAll(userRole, "\n", "")
  userRole = strings.ReplaceAll(userRole, "\r", "")

    // Check if the user has the required role
  return userRole == requiredRole
}

func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(data)
}
