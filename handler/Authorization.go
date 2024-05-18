package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func parseJWTFromRequest(r *http.Request) (*jwt.Token, error) {
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

	return token, err
}

func RequireRole(requiredRole string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := parseJWTFromRequest(r)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Extract the claims and check the role
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			role, ok := claims["role"].(string)
			if !ok || role != requiredRole {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
		} else {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// If the role is valid, call the next handler
		next.ServeHTTP(w, r)
	}
}

func ExtractUserIDFromJWT(r *http.Request) (uint, error) {
	token, err := parseJWTFromRequest(r)
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
