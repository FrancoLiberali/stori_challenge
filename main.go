package main

import (
	"flag"
	"log"
	"os"

	"github.com/FrancoLiberali/stori_challenge/app"
	"github.com/FrancoLiberali/stori_challenge/app/models"
)

const helpHeader = `Stori challenge by Franco Liberali.
Usage: stori_challenge [FLAGS]

FLAGS:
`

func main() {
	csvFileNameParam := flag.String("file", "", "CSV file with transactions to be processed")
	destinationEmail := flag.String("email", "", "Email address where the results of the process will be sent")
	help := flag.Bool("help", false, "Help")

	flag.Parse()

	if *help || *csvFileNameParam == "" || *destinationEmail == "" {
		print(helpHeader) //nolint:forbidigo // here is correct to print
		flag.PrintDefaults()

		os.Exit(1)
	}

	db, err := app.NewDBConnection()
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = db.AutoMigrate(
		models.User{},
		models.Transaction{},
	)
	if err != nil {
		log.Fatalln(err.Error())
	}

	processService, err := app.NewService(db)
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = processService.Process(*csvFileNameParam, *destinationEmail)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Transactions processed successfully")
}
