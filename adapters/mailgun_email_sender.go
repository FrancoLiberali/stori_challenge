package adapters

import (
	"context"
	"fmt"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

type MailgunEmailSender struct {
	APIKey string
}

const (
	mailgunDomain = "sandboxef42aaae33d8465e98a947fcce8bbd54.mailgun.org" // https://app.mailgun.com/app/domains
	senderEmail   = "stori-challenge@" + mailgunDomain
	sendTimeout   = time.Second * 10
)

// Send sends an email to recipient with the subject and body received by parameter.
//
// Returns ErrSendingEmail if an error is produced
func (sender MailgunEmailSender) Send(recipient, subject, body string) error {
	mg := mailgun.NewMailgun(mailgunDomain, sender.APIKey)

	message := mg.NewMessage(senderEmail, subject, body, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), sendTimeout)
	defer cancel()

	_, _, err := mg.Send(ctx, message)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrSendingEmail, err.Error())
	}

	return nil
}
