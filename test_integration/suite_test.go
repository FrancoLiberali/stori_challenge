package testintegration

import (
	"html/template"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	"github.com/FrancoLiberali/cql"
	"github.com/FrancoLiberali/stori_challenge/app"
	mocks "github.com/FrancoLiberali/stori_challenge/app/mocks/adapters"
	"github.com/FrancoLiberali/stori_challenge/app/models"
	"github.com/FrancoLiberali/stori_challenge/app/repository"
	"github.com/FrancoLiberali/stori_challenge/app/service"
)

type IntTestSuite struct {
	suite.Suite
	db *gorm.DB
}

func (ts *IntTestSuite) SetupTest() {
	ts.T().Setenv(app.EmailPublicAPIKeyEnvVar, "asd")
	ts.T().Setenv(app.EmailPrivateAPIKeyEnvVar, "asd")
	ts.T().Setenv(app.DBURLEnvVar, host)
	ts.T().Setenv(app.DBPortEnvVar, strconv.Itoa(port))
	ts.T().Setenv(app.DBUserEnvVar, username)
	ts.T().Setenv(app.DBPasswordEnvVar, password)
	ts.T().Setenv(app.DBNameEnvVar, dbName)
	ts.T().Setenv(app.DBSSLEnvVar, sslMode)
	CleanDB(ts.db)
}

func (ts *IntTestSuite) TestProcessLocalCSVFileSendsEmailWhenUserDoesNotExists() {
	mockEmailSender := mocks.NewEmailSender(ts.T())
	userRepository := repository.UserRepository{}

	storiService, err := app.NewService(ts.db)
	ts.Require().NoError(err)

	storiService.EmailService = service.EmailService{
		Template:    template.Must(template.ParseFiles("../app/html/email.html")),
		EmailSender: mockEmailSender,
	}

	mockEmailSender.On(
		"Send",
		"client@mail.com",
		"Stori transaction summary",
		mock.MatchedBy(func(emailBody string) bool {
			return strings.Contains(emailBody, "39.74") && strings.Contains(emailBody, "-15.38") && strings.Contains(emailBody, "35.25")
		}),
	).Return(nil)

	err = storiService.Process("../data/txns1.csv", "client@mail.com")
	ts.Require().NoError(err)

	// check that transactions are created
	transactionAmount, err := cql.Query[models.Transaction](
		ts.db,
	).Count()
	ts.Require().NoError(err)
	ts.Equal(int64(4), transactionAmount)

	// check that user balance is created
	user, err := userRepository.GetByEmail(ts.db, "client@mail.com")
	ts.Require().NoError(err)
	ts.Equal(decimal.NewFromFloat32(39.74), user.Balance)
}

func (ts *IntTestSuite) TestProcessLocalCSVFileSendsEmailWhenUserExists() {
	mockEmailSender := mocks.NewEmailSender(ts.T())
	userRepository := repository.UserRepository{}
	storiService, err := app.NewService(ts.db)
	ts.Require().NoError(err)

	storiService.EmailService = service.EmailService{
		Template:    template.Must(template.ParseFiles("../app/html/email.html")),
		EmailSender: mockEmailSender,
	}

	// create user
	ts.Require().NoError(userRepository.Save(ts.db, &models.User{Email: "client@mail.com", Balance: decimal.NewFromFloat(30.3)}))

	mockEmailSender.On(
		"Send",
		"client@mail.com",
		"Stori transaction summary",
		mock.MatchedBy(func(emailBody string) bool {
			return strings.Contains(emailBody, "70.04") &&
				strings.Contains(emailBody, "39.74") &&
				strings.Contains(emailBody, "-15.38") &&
				strings.Contains(emailBody, "35.25")
		}),
	).Return(nil)

	err = storiService.Process("../data/txns1.csv", "client@mail.com")
	ts.Require().NoError(err)

	// check that transactions are created
	transactionAmount, err := cql.Query[models.Transaction](
		ts.db,
	).Count()
	ts.Require().NoError(err)
	ts.Equal(int64(4), transactionAmount)

	// check that user balance is updated
	user, err := userRepository.GetByEmail(ts.db, "client@mail.com")
	ts.Require().NoError(err)
	ts.Equal(decimal.NewFromFloat32(39.74).Add(decimal.NewFromFloat32(30.3)), user.Balance)
}

func (ts *IntTestSuite) TestProcessS3CSVFileSendsEmail() {
	mockEmailSender := mocks.NewEmailSender(ts.T())
	mockS3Reader := mocks.NewCSVReader(ts.T())
	storiService, err := app.NewService(ts.db)
	ts.Require().NoError(err)

	storiService.EmailService = service.EmailService{
		Template:    template.Must(template.ParseFiles("../app/html/email.html")),
		EmailSender: mockEmailSender,
	}

	storiService.TransactionsReader = service.TransactionsReader{
		S3CSVReader: mockS3Reader,
	}

	mockS3Reader.On("Read", "fl-stori-challenge/txns1.csv").Return(
		[][]string{
			{"0", "7/15", "+60.5"},
			{"1", "7/28", "-10.3"},
			{"2", "8/2", "-20.46"},
			{"3", "8/13", "+10"},
		}, nil,
	)
	mockEmailSender.On(
		"Send",
		"client@mail.com",
		"Stori transaction summary",
		mock.MatchedBy(func(emailBody string) bool {
			return strings.Contains(emailBody, "39.74") && strings.Contains(emailBody, "-15.38") && strings.Contains(emailBody, "35.25")
		}),
	).Return(nil)

	err = storiService.Process("s3://fl-stori-challenge/txns1.csv", "client@mail.com")
	ts.Require().NoError(err)
}
