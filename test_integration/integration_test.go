package testintegration

import (
	"log"
	"strconv"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/FrancoLiberali/stori_challenge/app"
)

const (
	host     = "localhost"
	port     = 5432
	username = "stori"
	password = "stori_challenge2024"
	sslMode  = "disable"
	dbName   = "stori_db"
)

func TestMain(t *testing.T) {
	t.Setenv(app.EmailPublicAPIKeyEnvVar, "asd")
	t.Setenv(app.EmailPrivateAPIKeyEnvVar, "asd")
	t.Setenv(app.DBURLEnvVar, host)
	t.Setenv(app.DBPortEnvVar, strconv.Itoa(port))
	t.Setenv(app.DBUserEnvVar, username)
	t.Setenv(app.DBPasswordEnvVar, password)
	t.Setenv(app.DBNameEnvVar, dbName)
	t.Setenv(app.DBSSLEnvVar, sslMode)

	db, err := app.NewDBConnection()
	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(ListOfTables...)
	if err != nil {
		log.Fatalln(err)
	}

	suite.Run(t, &IntTestSuite{db: db})
}
