package testintegration

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/FrancoLiberali/cql/logger"
)

const (
	username = "stori"
	password = "stori_challenge2024"
	host     = "localhost"
	port     = 5432
	sslMode  = "disable"
	dbName   = "stori_db"
)

func TestMain(t *testing.T) {
	db, err := NewDBConnection()
	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(ListOfTables...)
	if err != nil {
		log.Fatalln(err)
	}

	suite.Run(t, &IntTestSuite{db: db})
}

func NewDBConnection() (*gorm.DB, error) {
	dialector := postgres.Open(
		fmt.Sprintf(
			"user=%s password=%s host=%s port=%d sslmode=%s dbname=%s",
			username, password, host, port, sslMode, dbName,
		),
	)

	return OpenWithRetry(
		dialector,
		logger.Default.ToLogMode(logger.Info),
		10, time.Duration(5)*time.Second,
	)
}
