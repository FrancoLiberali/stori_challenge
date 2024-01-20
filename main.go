package main

import (
	"flag"
	"log"
	"os"

	"github.com/FrancoLiberali/stori_challenge/adapters"
	"github.com/FrancoLiberali/stori_challenge/service"
)

//go:generate mockery --all --keeptree

const helpHeader string = `Stori challenge by Franco Liberali.
Usage: go run . [FLAGS]

FLAGS:
`

const MailgunApiKeyEnvVar = "MAILGUN_API_KEY"

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

	mailgunApiKey := os.Getenv(MailgunApiKeyEnvVar)
	if mailgunApiKey == "" {
		print("Mailgun api key env variable not configured") //nolint:forbidigo // here is correct to print
		os.Exit(1)
	}

	processService := service.Service{
		CSVReader: adapters.LocalCsvReader{},
		EmailSender: adapters.MailgunEmailSender{
			ApiKey: mailgunApiKey,
		},
	}

	err := processService.Process(*csvFileName, *destinationEmail)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
