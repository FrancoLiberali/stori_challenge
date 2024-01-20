package main

import (
	"flag"
	"log"
	"os"

	"github.com/FrancoLiberali/stori_challenge/app/adapters"
	"github.com/FrancoLiberali/stori_challenge/app/service"
)

//go:generate mockery --all --keeptree

const helpHeader string = `Stori challenge by Franco Liberali.
Usage: stori_challenge [FLAGS]

FLAGS:
`

const (
	emailPublicAPIKeyEnvVar  = "EMAIL_PUBLIC_API_KEY"  //nolint:gosec // just the env var name
	emailPrivateAPIKeyEnvVar = "EMAIL_PRIVATE_API_KEY" //nolint:gosec // just the env var name
)

func main() {
	csvFileName := flag.String("file", "", "CSV file with transactions to be processed")
	destinationEmail := flag.String("email", "", "Email address where the results of the process will be sent")
	help := flag.Bool("help", false, "Help")

	flag.Parse()

	if *help || *csvFileName == "" || *destinationEmail == "" {
		print(helpHeader) //nolint:forbidigo // here is correct to print
		flag.PrintDefaults()

		os.Exit(1)
	}

	emailPublicAPIKey := os.Getenv(emailPublicAPIKeyEnvVar)
	emailPrivateAPIKey := os.Getenv(emailPrivateAPIKeyEnvVar)

	if emailPublicAPIKey == "" || emailPrivateAPIKey == "" {
		print("Email api key env variables not configured") //nolint:forbidigo // here is correct to print
		os.Exit(1)
	}

	processService := service.Service{
		CSVReader: adapters.LocalCsvReader{},
		EmailSender: adapters.MailJetEmailSender{
			PublicAPIKey:  emailPublicAPIKey,
			PrivateAPIKey: emailPrivateAPIKey,
		},
	}

	err := processService.Process(*csvFileName, *destinationEmail)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
