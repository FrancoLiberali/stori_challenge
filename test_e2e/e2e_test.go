package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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
	csvFileName    = "csv_file.csv"
	s3BucketRegion = "us-east-2"
	s3Protocol     = "s3://"
	s3BucketName   = "fl-stori-challenge"
	appName        = "../process.sh"
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
	sc.Step(`^there is a S3 CSV file with following data$`, s3CSVFile)
	sc.Step(`^there is a local CSV file with following data$`, localCSVFile)
	sc.Step(`^the system is executed$`, executeSystem)
	sc.Step(`^I receive an email with subject "([^"]*)" and with the following information$`, iReceiveTheEmail)
}

// Creates in local a CSV file called csv_file.csv with the content of the godog.Table
func localCSVFile(fileContent *godog.Table) error {
	csvFile, err := os.Create(csvFileName)
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

	fileName = csvFileName

	return nil
}

// Creates in s3 a CSV file called csv_file.csv with the content of the godog.Table
func s3CSVFile(fileContent *godog.Table) error {
	err := localCSVFile(fileContent)
	if err != nil {
		return err
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(s3BucketRegion),
	})
	if err != nil {
		return err
	}

	svc := s3.New(sess)

	file, err := os.Open(csvFileName)
	if err != nil {
		return err
	}

	defer file.Close()

	// Read the contents of the file into a buffer
	var buf bytes.Buffer
	if _, err = io.Copy(&buf, file); err != nil {
		return err
	}

	acl := "public-read"

	// Upload the contents of the buffer to S3
	_, err = svc.PutObject(&s3.PutObjectInput{
		ACL:    &acl,
		Bucket: aws.String(s3BucketName),
		Key:    aws.String(csvFileName),
		Body:   bytes.NewReader(buf.Bytes()),
	})
	if err != nil {
		return err
	}

	fileName = s3Protocol + s3BucketName + "/" + csvFileName

	return nil
}

// Executes the transaction processing system
func executeSystem() error {
	output, err := exec.Command(appName, fileName, receiveInbox.EmailAddress).CombinedOutput()
	if err != nil {
		log.Println(string(output))
	}

	return err
}

// Checks that the receiveInbox got an email with the information from the godog.Table
func iReceiveTheEmail(subject string, content *godog.Table) error {
	// fetch the email for receiveInbox
	email, err := getLastEmail(mailCtx, mailClient, receiveInbox.Id)
	if err != nil {
		return err
	}

	// assert email content
	err = assertExpectedAndActual(assert.Equal, subject, *email.Subject)
	if err != nil {
		return err
	}

	for _, row := range content.Rows {
		var contains string

		value := row.Cells[1].Value

		switch row.Cells[0].Value {
		case "Total balance":
			contains = fmt.Sprintf("<span>Total balance: <strong>$%s</strong></span>", value)
		case "Average credit":
			contains = fmt.Sprintf(`<span class="avg-cre"><span><span><strong>$%s</strong></span></span></span>`, value)
		case "Average debit":
			contains = fmt.Sprintf(`<span class="avg-de"><span><span><strong>$%s</strong></span></span></span>`, value)
		default:
			month := strings.TrimPrefix(row.Cells[0].Value, "Number of transactions in ")
			contains = fmt.Sprintf(`<span class="n-%s"><span><span><strong>%s</strong></span></span></span>`, month, value)
		}

		err = assertExpectedAndActual(assert.Contains, *email.Body, contains)
		if err != nil {
			return err
		}
	}

	return nil
}
