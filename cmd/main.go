package main

import (
	"log/slog"
	"notificationService/internal/application"
)

// main starts the notification service in developer mode
// After starting, notifications can be sent to http://localhost:3000/api/v1/notifications
func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	app := application.New()
	err := app.Run()
	if err != nil {
		slog.Error("Failed to run the app", "error", err)
	}
}
