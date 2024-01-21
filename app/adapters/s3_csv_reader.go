package adapters

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3CSVReader struct {
	LocalCSVReader LocalCSVReader
}

const (
	awsRegion      = "us-east-2"
	s3FileNameSize = 2 // bucket/filename
)

// Read reads a CSV file by its name.
//
// Returns the list of rows of the CSV file
// or ErrReadingFile if an error is produced
func (reader S3CSVReader) Read(fileName string) ([][]string, error) {
	fileNameSplitted := strings.Split(fileName, string(filepath.Separator))
	if len(fileNameSplitted) != s3FileNameSize {
		return nil, errReadingFile(fileName, "invalid s3 path")
	}

	bucket := fileNameSplitted[0]
	item := fileNameSplitted[1]

	itemFileName := filepath.Join(os.TempDir(), item)

	file, err := os.Create(itemFileName)
	if err != nil {
		return nil, errReadingFile(fileName, err.Error())
	}

	defer file.Close()

	// Initialize a session in us-west-2
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.AnonymousCredentials,
	})
	if err != nil {
		return nil, errReadingFile(fileName, err.Error())
	}

	downloader := s3manager.NewDownloader(sess)

	_, err = downloader.Download(
		file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(item),
		})
	if err != nil {
		return nil, errReadingFile(fileName, fmt.Sprintf("Unable to download item: %s", err.Error()))
	}

	return reader.LocalCSVReader.Read(itemFileName)
}

func errReadingFile(fileName string, internalErr string) error {
	return fmt.Errorf("%w %s: %s", ErrReadingFile, fileName, internalErr)
}
