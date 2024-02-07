package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var notification = Notification{
	To:      "1234567890",
	Message: "Hello, World!",
}

func TestGetSMSNotifier(t *testing.T) {
	sms, err := GetNotifier(SMS)
	assert.Nil(t, err)
	assert.IsType(t, &SMSNotifier{}, sms)
	msg := sms.Send(notification)
	assert.Contains(t, msg, "was sent to number")
}

func TestGetEmailNotifier(t *testing.T) {
	email, err := GetNotifier(Email)
	assert.Nil(t, err)
	assert.IsType(t, &EmailNotifier{}, email)

	msg := email.Send(notification)
	assert.Contains(t, msg, "was sent to email")
}

func TestGetUnkownNotifier(t *testing.T) {
	_, err := GetNotifier(3)
	assert.NotNil(t, err)
	assert.Equal(t, "Notifier with id 3 not found", err.Error())
}
