package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"notificationService/internal/messaging"
	"strings"
)

// UseNotificationsRoute adds the api/v1/notification POST endpoint to the router
func UseNotificationsRoute(router *chi.Mux, receiverInstances []messaging.Receiver) {
	router.Post("/api/v1/notifications", func(w http.ResponseWriter, r *http.Request) {
		notification, err := NewNotification(r)
		if err != nil {
			RespondWithBadRequestResult(w, fmt.Sprintf("The request payload is not valid json. Error: %s", err.Error()))
			return
		}

		err = ValidateNotification(*notification)
		if err != nil {
			RespondWithBadRequestResult(w, fmt.Sprintf("The notification payload is invalid. Error: %s", err.Error()))
			return
		}

		go messaging.DispatchAsync(receiverInstances, *notification)

		RespondWithAcceptedResult(w)
	})
}

func RespondWithAcceptedResult(w http.ResponseWriter) {
	responseBody := struct {
		Message string `json:"message"`
	}{
		Message: "Ok",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(responseBody)
}

func RespondWithBadRequestResult(w http.ResponseWriter, message string) {
	responseBody := struct {
		Message string `json:"message"`
	}{
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(responseBody)
}

func NewNotification(r *http.Request) (*messaging.Notification, error) {
	var notification messaging.Notification
	err := json.NewDecoder(r.Body).Decode(&notification)
	if err != nil {
		return nil, err
	}

	// TODO check further things. E.g. is https://www.invicti.com/learn/json-injection/ a thing in golang?

	return &notification, nil
}

// ValidateNotification checks that the content of the notification complies with the specifications
func ValidateNotification(notification messaging.Notification) error {
	title := strings.TrimSpace(notification.Title)
	if title == "" {
		return errors.New("the 'title' property is required and should not be blank")
	}

	description := strings.TrimSpace(notification.Description)
	if description == "" {
		return errors.New("the 'description' property is required and should not be blank")
	}

	typeName := notification.Type
	switch typeName {
	case "Info":
	case "Warning":
	case "Error":
	case "Debug":
	case "Fatal":
		break
	default:
		return fmt.Errorf("the 'type' property with value '%s' is invalid. It should be one of the following values: Info, Warning, Error, Debug, Fatal", typeName)
	}

	return nil
}
