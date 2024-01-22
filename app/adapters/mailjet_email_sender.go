package adapters

import (
	"fmt"

	"github.com/mailjet/mailjet-apiv3-go/v4"
)

type MailJetEmailSender struct {
	PublicAPIKey  string
	PrivateAPIKey string
}

const (
	senderEmail = "franco.liberali@gmail.com"
	senderName  = "Franco Liberali"
)

// Send sends an email to recipient with the subject and body received by parameter.
//
// Returns ErrSendingEmail if an error is produced
func (sender MailJetEmailSender) Send(recipient, subject, body string) error {
	client := mailjet.NewMailjetClient(sender.PublicAPIKey, sender.PrivateAPIKey)

	messages := mailjet.MessagesV31{Info: []mailjet.InfoMessagesV31{
		{
			From:     &mailjet.RecipientV31{Email: senderEmail, Name: senderName},
			To:       &mailjet.RecipientsV31{mailjet.RecipientV31{Email: recipient}},
			Subject:  subject,
			HTMLPart: body,
		},
	}}

	_, err := client.SendMailV31(&messages)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrSendingEmail, err.Error())
	}

	return nil
}
