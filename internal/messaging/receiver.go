package messaging

// Receiver interface must be implemented by concrete receivers
//
// Note: In case a receiver failed to send the notification to the receiver system,
// the receiver must handle the failure itself.
type Receiver interface {
	Receive(notification Notification) error
}
