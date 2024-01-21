package main

import (
	"context"
	"errors"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/FrancoLiberali/stori_challenge/app"
)

type ProcessTransactionsEvent struct {
	FileName string `json:"file"`
	Email    string `json:"email"`
}

var ErrNilEvent = errors.New("received nil event")

func HandleRequest(_ context.Context, event *ProcessTransactionsEvent) (*string, error) {
	if event == nil {
		return nil, ErrNilEvent
	}

	processService, err := app.NewService()
	if err != nil {
		return nil, err
	}

	err = processService.Process(event.FileName, event.Email)
	if err != nil {
		return nil, err
	}

	successMessage := "Process executed successfully"

	return &successMessage, nil
}

func main() {
	lambda.Start(HandleRequest)
}
