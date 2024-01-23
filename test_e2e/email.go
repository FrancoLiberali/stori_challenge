package main

import (
	"context"
	"net/http"

	"github.com/antihax/optional"
	mailslurp "github.com/mailslurp/mailslurp-client-go"
	"github.com/stretchr/testify/assert"
)

const (
	mailSlurpAPIKey  = "02c92c8842430fccb6f415a1f9eacfc93a9f49cb0552071272dcc3cab8214897" //nolint:gosec // simplification
	receiveInboxID   = "76b6f705-49a1-4c96-b42c-fa75b122bc90"
	mailSlurpTimeout = 30000 // miliseconds
)

// Returns a client and a context that can be used to communicate with MailSlurp
func getMailSlurpClient() (*mailslurp.APIClient, context.Context) {
	// create a context with your api key
	ctx := context.WithValue(context.Background(), mailslurp.ContextAPIKey, mailslurp.APIKey{Key: mailSlurpAPIKey})

	// create mailslurp client
	config := mailslurp.NewConfiguration()
	client := mailslurp.NewAPIClient(config)

	return client, ctx
}

// Returns the MailSlurp inbox where emails will be received
func getReceiveInbox(ctx context.Context, client *mailslurp.APIClient) (mailslurp.InboxDto, error) {
	return getInbox(ctx, client, receiveInboxID)
}

// Returns a MailSlurp inbox from its id
func getInbox(ctx context.Context, client *mailslurp.APIClient, id string) (mailslurp.InboxDto, error) {
	inbox, res, err := client.InboxControllerApi.GetInbox(ctx, id)

	return inbox, assertResponseAndError(res, err)
}

// Returns last email from a MailSlurp inbox
func getLastEmail(ctx context.Context, client *mailslurp.APIClient, inboxID string) (mailslurp.Email, error) {
	waitOpts := &mailslurp.WaitForLatestEmailOpts{
		InboxId:    optional.NewInterface(inboxID),
		Timeout:    optional.NewInt64(mailSlurpTimeout),
		UnreadOnly: optional.NewBool(true),
	}

	email, res, err := client.WaitForControllerApi.WaitForLatestEmail(ctx, waitOpts)

	return email, assertResponseAndError(res, err)
}

// Assert that the response code is 200 OK
func assertResponseAndError(res *http.Response, err error) error {
	defer res.Body.Close()

	err = assertError(assert.NoError, err)
	if err != nil {
		return err
	}

	return assertExpectedAndActual(assert.Equal, http.StatusOK, res.StatusCode)
}
