package routes

import (
	"notificationService/internal/messaging"
	"testing"
)

func TestValidateNotification(t *testing.T) {
	type args struct {
		notification messaging.Notification
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"title is required", args{messaging.Notification{Title: "", Description: "A description", Type: "Info"}}, true},
		{"title should not be blank", args{messaging.Notification{Title: " ", Description: "A description", Type: "Info"}}, true},
		{"description is required", args{messaging.Notification{Title: "A title", Description: "", Type: "Info"}}, true},
		{"description should not be blank", args{messaging.Notification{Title: "A title", Description: " ", Type: "Info"}}, true},

		{"type Info is valid", args{messaging.Notification{Title: "A title", Description: "A description", Type: "Info"}}, false},
		{"type Warning is valid", args{messaging.Notification{Title: "A title", Description: "A description", Type: "Warning"}}, false},
		{"type Fatal is valid", args{messaging.Notification{Title: "A title", Description: "A description", Type: "Fatal"}}, false},
		{"type Debug is valid", args{messaging.Notification{Title: "A title", Description: "A description", Type: "Debug"}}, false},
		{"type Error is valid", args{messaging.Notification{Title: "A title", Description: "A description", Type: "Error"}}, false},

		{"type should be case sensitive", args{messaging.Notification{Title: "A title", Description: "A description ", Type: "info"}}, true},
		{"type should consider blanks", args{messaging.Notification{Title: "A title", Description: "A description ", Type: " info "}}, true},

		// TODO Add more edge test cases...
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateNotification(tt.args.notification); (err != nil) != tt.wantErr {
				t.Errorf("ValidateNotification() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
