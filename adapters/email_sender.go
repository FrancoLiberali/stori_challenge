package adapters

import (
	"errors"
)

type EmailSender interface {
	// Send sends an email to recipient with the subject and body received by parameter.
	//
	// Returns ErrSendingEmail if an error is produced
	Send(recipient, subject, body string) error
}

var ErrSendingEmail = errors.New("error while sending email")
