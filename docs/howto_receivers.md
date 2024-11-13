# How to add or modify concrete receivers

Concrete receivers (CR) act like proxies for actual Receiver System. For instance, the `GithubIssueReceiver` sends 
notifications of type warning to github as new issues. Open `\internal\receivers\githubIssueReceiver.go` 
to see how a concrete receiver looks like. 

## Implementing a CR from scratch
Implementing a concrete receiver is straight forward.

1. Create a new *receiver.go file in folder `/internal/receivers`. E.g. `twilioReceiver.go`
2. Create a type `TwilioReceiver ` and implement the interface `Receiver`  
Your type should look like this
``` golang 
package receivers

import (
	"notificationService/internal/messaging"
)

type TwilioReceiver struct {
}

func NewTwilioReceiver() *TwilioReceiver {
	return &TwilioReceiver{}
}

func (instance *TwilioReceiver) Receive(notification messaging.Notification) error {
	shouldSend := notification.Type == "Fatal"
	if !shouldSend {
		slog.Debug("Ignore unsupported notification type", "notification", notification)
		return nil
	}
	
    //  Implement Twilio integration here
}
```

The `Receive` method takes a `Notification` instance as argument that is passed in from the dispatcher whenever 
a sender system sends a notification to the notification service.  

3. Add the code to send the notification to Twilio in the `Receive` method.
4. In order to get your receiver registered on start-up, open the `receivers.go` file (same folder) and add an instance of
your receiver in the `NewReceivers` function.

> [!NOTE]
> All notifications will be sent to your receiver. If you want to ignore notification then just don't handle them.

## Adding unit and integration tests

// TODO