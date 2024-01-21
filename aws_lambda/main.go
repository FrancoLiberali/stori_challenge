package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/FrancoLiberali/stori_challenge/app"
)

const (
	fileJSONKey  = "file"
	emailJSONKey = "email"
)

var (
	ErrNilEvent   = errors.New("received nil event")
	ErrBadRequest = errors.New("bad request")
)

func HandleRequest(_ context.Context, event *events.APIGatewayV2HTTPRequest) (*string, error) {
	if event == nil {
		return nil, ErrNilEvent
	}

	body := map[string]string{}

	err := json.Unmarshal([]byte(event.Body), &body)
	if err != nil {
		return nil, err
	}

	fileName, isPresent := body[fileJSONKey]
	if !isPresent {
		return nil, badRequest(fileJSONKey)
	}

	email, isPresent := body[emailJSONKey]
	if !isPresent {
		return nil, badRequest(emailJSONKey)
	}

	processService, err := app.NewService()
	if err != nil {
		return nil, err
	}

	err = processService.Process(fileName, email)
	if err != nil {
		return nil, err
	}

	successMessage := "Process executed successfully"

	return &successMessage, nil
}

func badRequest(paramName string) error {
	return fmt.Errorf("%w: excepted param %s", ErrBadRequest, paramName)
}

func main() {
	lambda.Start(HandleRequest)
}
