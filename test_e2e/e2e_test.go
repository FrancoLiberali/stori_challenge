package main

import (
	"context"
	"encoding/csv"
	"os"
	"os/exec"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	messages "github.com/cucumber/messages/go/v21"
	"github.com/elliotchance/pie/v2"
	mailslurp "github.com/mailslurp/mailslurp-client-go"
	"github.com/stretchr/testify/assert"
)

var (
	mailClient   *mailslurp.APIClient
	mailCtx      context.Context
	receiveInbox mailslurp.InboxDto
	fileName     string
)

const (
	localFileName = "local_csv_file.csv"
)

func init() {
	opts := godog.Options{Output: colors.Colored(os.Stdout)}
	godog.BindCommandLineFlags("godog.", &opts)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	// create a MailSlurp client and context with api key
	mailClient, mailCtx = getMailSlurpClient()

	// get receive inbox for testing
	var err error
	receiveInbox, err = getReceiveInbox(mailCtx, mailClient)

	if err != nil {
		t.Fatal("Error getting receive inbox:", err.Error())
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(sc *godog.ScenarioContext) {
	sc.Step(`^there is a local CSV file with following data$`, localCSVFile)
	sc.Step(`^the system is executed$`, executeSystem)
	sc.Step(`^I receive an email with subject "([^"]*)" and with the following information$`, iReceiveTheEmail)
}

// Creates a CSV file called local_csv_file.csv with the content of the godog.Table
func localCSVFile(fileContent *godog.Table) error {
	csvFile, err := os.Create(localFileName)
	if err != nil {
		return err
	}

	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)

	for _, row := range fileContent.Rows {
		err = csvWriter.Write(pie.Map(row.Cells, func(cell *messages.PickleTableCell) string {
			return cell.Value
		}))
		if err != nil {
			return err
		}
	}

	csvWriter.Flush()

	fileName = localFileName

	return nil
}

// Executes the transaction processing system
func executeSystem() error {
	app := "stori_challenge"
	argFile := "-file"
	argEmail := "-email"

	return exec.Command(app, argFile, fileName, argEmail, receiveInbox.EmailAddress).Run()
}

// Checks that the receiveInbox got an email with the information from the godog.Table
func iReceiveTheEmail(subject string, content *godog.Table) error {
	// fetch the email for receiveInbox
	email, err := getLastEmail(mailCtx, mailClient, receiveInbox.Id)
	if err != nil {
		return err
	}

	// assert email content
	err = assertExpectedAndActual(assert.Equal, subject, email.Subject)
	if err != nil {
		return err
	}

	for _, row := range content.Rows {
		err = assertExpectedAndActual(assert.Contains, row.Cells[0].Value, email.Body)
		if err != nil {
			return err
		}
	}

	return nil
}
