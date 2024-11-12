package tests

import (
	"bytes"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"net/http"
	"notificationService/internal/application"
	"notificationService/internal/messaging"
	"notificationService/internal/routes"
	"testing"
	"time"
)

// Test_Integration starts the notification service locally to run integration tests
// against a concrete receiver (stub)
func Test_Integration(t *testing.T) {
	testReceiver := StubReceiver{}
	notificationUrl := "http://localhost:3000/api/v1/notifications"
	StartNotificationService(&testReceiver)

	time.Sleep(1 * time.Second)

	t.Run("send notification to receiver with valid message", func(t *testing.T) {
		testReceiver.ReceivedNotification = nil
		jsonValue, _ := json.Marshal(messaging.Notification{
			Title:       "A title",
			Description: "A description",
			Type:        "Warning",
		})
		resp, err := http.Post(notificationUrl, "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		AssertResponseStatusAndMessage(t, resp, "202 Accepted", "Ok")
		AssertReceivedNotification(t, testReceiver)
	})

	t.Run("send notification to receiver with empty title", func(t *testing.T) {
		testReceiver.ReceivedNotification = nil
		jsonValue, _ := json.Marshal(messaging.Notification{
			Title:       "",
			Description: "A description",
			Type:        "Warning",
		})
		resp, err := http.Post(notificationUrl, "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		AssertResponseStatusAndMessage(t, resp, "400 Bad Request", "The notification payload is invalid. Error: the 'title' property is required and should not be blank")
	})
}

func StartNotificationService(testReceiver *StubReceiver) {
	app := application.New()
	go func() {
		err := app.UseAndRun([]func(router *chi.Mux){
			func(router *chi.Mux) {
				routes.UseNotificationsRoute(router, []messaging.Receiver{
					testReceiver,
				})
			},
		})
		if err != nil {
			slog.Error(err.Error())
		}
	}()
}

func AssertResponseStatusAndMessage(t *testing.T, resp *http.Response, status string, message string) {
	body := struct {
		Message string `json:"message"`
	}{}
	json.NewDecoder(resp.Body).Decode(&body)
	assert.Equal(t, body.Message, message)
	assert.Equal(t, status, resp.Status)
}

func AssertReceivedNotification(t *testing.T, testReceiver StubReceiver) {
	assert.NotNil(t, testReceiver.ReceivedNotification)
	assert.Equal(t, testReceiver.ReceivedNotification.Title, "A title")
	assert.Equal(t, testReceiver.ReceivedNotification.Description, "A description")
	assert.Equal(t, testReceiver.ReceivedNotification.Type, "Warning")
}
