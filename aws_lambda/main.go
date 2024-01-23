package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"gorm.io/gorm"

	"github.com/FrancoLiberali/stori_challenge/app"
	"github.com/FrancoLiberali/stori_challenge/app/service"
)

const (
	fileJSONKey  = "file"
	emailJSONKey = "email"
)

var (
	ErrNilEvent   = errors.New("received nil event")
	ErrBadRequest = errors.New("bad request")
)

var (
	db             *gorm.DB
	processService *service.Service
)

func init() {
	var err error

	db, err = app.NewDBConnection()
	if err != nil {
		log.Fatalln(err.Error())
	}

	processService, err = app.NewService(db)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

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
