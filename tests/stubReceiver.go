package tests

import "notificationService/internal/messaging"

type StubReceiver struct {
	ReceivedNotification *messaging.Notification
}

func (instance *StubReceiver) Receive(notification messaging.Notification) error {
	instance.ReceivedNotification = &notification
	return nil
}
