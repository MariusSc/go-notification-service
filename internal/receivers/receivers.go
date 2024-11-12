package receivers

import (
	"notificationService/internal/messaging"
)

// NewReceivers creates a list of concrete notification receivers
func NewReceivers() []messaging.Receiver {
	return []messaging.Receiver{
		NewGithubIssueReceiverFromEnvVars(),
		// Add new receivers here
	}
}
