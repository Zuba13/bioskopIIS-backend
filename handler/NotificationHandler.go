package handler

import (
	"net/http"

	"bioskop.com/projekat/bioskopIIS-backend/service"
	"github.com/gorilla/mux"
)

type NotificationHandler struct {
	NotificationService service.NotificationService
}

func NewNotificationHandler(notificationService service.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		NotificationService: notificationService,
	}
}

func (nh *NotificationHandler) RegisterNotificationHandler(r *mux.Router) {
	r.HandleFunc("/notification/canceled-projection", RequireRole("user", nh.GetCanceledProjectionNotifications)).Methods("GET")
}

func (nh *NotificationHandler) GetCanceledProjectionNotifications(w http.ResponseWriter, r *http.Request) {
	userId, err := ExtractUserIDFromJWT(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	notifications, err := nh.NotificationService.GetCanceledProjectionNotifications(userId)
	if err != nil {
		customErr, ok := err.(*service.CustomError)
		if ok {
			http.Error(w, customErr.Message, service.ErrorCodeToHTTPStatus(customErr.Code))
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	respondWithJSON(w, http.StatusOK, notifications)
}
