package receivers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"notificationService/internal/messaging"
	"testing"
	"time"
)

func TestGithubIssueReceiver_CanSend(t *testing.T) {
	type args struct {
		message messaging.Notification
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"type Warning can be sent", args{messaging.Notification{"A title", "A description", "Warning"}}, true},
		{"type Info cannot be sent", args{messaging.Notification{"A title", "A description", "Info"}}, false},
		{"type Error cannot be sent", args{messaging.Notification{"A title", "A description", "Error"}}, false},
		{"type Debug cannot be sent", args{messaging.Notification{"A title", "A description", "Debug"}}, false},
		{"type Fatal cannot be sent", args{messaging.Notification{"A title", "A description", "Fatal"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			instance := &GithubIssueReceiver{}
			assert.Equalf(t, tt.want, instance.CanSend(tt.args.message), "CanSend(%v)", tt.args.message)
		})
	}
}

func TestGithubIssueReceiver_CreateIssue(t *testing.T) {
	type args struct {
		notification messaging.Notification
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{"issue is created", args{notification: messaging.Notification{
			Title:       fmt.Sprintf("Backup Failure (%v)", time.Now().Format(time.RFC822)),
			Description: "The backup failed due to a network problem",
			Type:        "Warning",
		}}, assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			instance := NewGithubIssueReceiverFromEnvVars()
			tt.wantErr(t, instance.CreateIssue(tt.args.notification), fmt.Sprintf("CreateIssue(%v)", tt.args.notification))
		})
	}
}
