package factory

import (
	"errors"
	"fmt"
)

const (
	SMS   = 1
	Email = 2
)

type Notification struct {
	To      string
	Message string
}

type Notifier interface {
	Send(Notification) string
}

type SMSNotifier struct{}

func (SMSNotifier) Send(n Notification) string {
	return fmt.Sprintf("%s was sent to number %s", n.Message, n.To)
}

type EmailNotifier struct{}

func (EmailNotifier) Send(n Notification) string {
	return fmt.Sprintf("%s was sent to email %s", n.Message, n.To)
}

func GetNotifier(n int) (Notifier, error) {
	switch n {
	case SMS:
		return &SMSNotifier{}, nil
	case Email:
		return &EmailNotifier{}, nil
	default:
		return nil, errors.New(fmt.Sprintf("Notifier with id %d not found", n))
	}
}
