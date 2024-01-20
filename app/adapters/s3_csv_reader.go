package adapters

import (
	"fmt"
	"os"
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
	fileNameSplitted := strings.Split(fileName, "/")
	if len(fileNameSplitted) != s3FileNameSize {
		return nil, fmt.Errorf("%w %s: %s", ErrReadingFile, fileName, "invalid s3 path")
	}

	bucket := fileNameSplitted[0]
	item := fileNameSplitted[1]

	file, err := os.Create(item)
	if err != nil {
		return nil, fmt.Errorf("%w %s: %s", ErrReadingFile, fileName, err.Error())
	}

	defer file.Close()
	defer os.Remove(item)

	// Initialize a session in us-west-2
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.AnonymousCredentials,
	})
	if err != nil {
		return nil, fmt.Errorf("%w %s: %s", ErrReadingFile, fileName, err.Error())
	}

	downloader := s3manager.NewDownloader(sess)

	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(item),
		})
	if err != nil {
		return nil, fmt.Errorf("%w %s: Unable to download item: %s", ErrReadingFile, fileName, err.Error())
	}

	return reader.LocalCSVReader.Read(item)
}
