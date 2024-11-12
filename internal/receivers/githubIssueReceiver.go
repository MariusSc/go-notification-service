package receivers

import (
	"context"
	"github.com/google/go-github/v66/github"
	"log/slog"
	"notificationService/internal/messaging"
	"os"
)

type GithubIssueReceiver struct {
	client *github.Client
	owner  string
	repo   string
}

// NewGithubIssueReceiverFromEnvVars creates a new instance based on environment variables
//
// Note: The following env vars needs to be set
// - GITHUB_TOKEN
// - GITHUB_OWNER
// - GITHUB_REPO
func NewGithubIssueReceiverFromEnvVars() *GithubIssueReceiver {
	client := github.NewClient(nil).WithAuthToken(os.Getenv("GITHUB_TOKEN"))
	return &GithubIssueReceiver{
		client: client,
		owner:  os.Getenv("GITHUB_OWNER"),
		repo:   os.Getenv("GITHUB_REPO"),
	}
}

// Receive creates new issues in Github based on the notification message
func (instance *GithubIssueReceiver) Receive(notification messaging.Notification) error {
	if !instance.CanSend(notification) {
		// It is okay to return nil, as it isn't an error if the notification should not be handled
		return nil
	}

	err := instance.CreateIssue(notification)
	if err != nil {
		return err
	}

	return nil
}

func (instance *GithubIssueReceiver) CanSend(notification messaging.Notification) bool {
	isSupportedType := notification.Type == "Warning"
	if !isSupportedType {
		slog.Debug("Not supported notification type", "notification", notification)
		return false
	}

	return true
}

func (instance *GithubIssueReceiver) CreateIssue(notification messaging.Notification) error {
	issueRequest := &github.IssueRequest{
		Title:       &notification.Title,
		Body:        &notification.Description,
		Labels:      nil,
		Assignee:    nil,
		State:       nil,
		StateReason: nil,
		Milestone:   nil,
		Assignees:   nil,
	}

	ctx := context.Background()
	_, _, err := instance.client.Issues.Create(ctx, instance.owner, instance.repo, issueRequest)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}
