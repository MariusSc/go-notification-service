package messaging

import "log/slog"

// DispatchAsync loops over all receivers and calls the Receive method with the given notification
//
// Note: As of the time being, no error handling or retry logic is implemented.
// That means, if a receiver fails to send the notification it will be ignored.
func DispatchAsync(receivers []Receiver, message Notification) {
	for _, receiver := range receivers {
		go DispatchInternal(receiver, message)
	}
}

func DispatchInternal(receiver Receiver, message Notification) {
	err := receiver.Receive(message)
	if err != nil {
		// TODO implement a retry-mechanism
		// TODO add to Dead-Letter-QUEUE when implementing a queueing or a pub/sub mechanism
		slog.Info("Error while dispatching message", "error", err.Error())
		return
	}
}
