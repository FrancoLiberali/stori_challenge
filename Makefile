install_dependencies:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/vektra/mockery/v2@v2.40.1

lint:
	golangci-lint run
	cd aws_lambda && golangci-lint run --config ../.golangci.yml
	cd app && golangci-lint run --config ../.golangci.yml
	cd test_integration && golangci-lint run --config ../.golangci.yml
	cd test_e2e && golangci-lint run --config ../.golangci.yml

test_unit:
	go test -v ./app/...

test_integration:
	go test -v ./test_integration

test_e2e:
	go test -v -count=1 ./test_e2e

aws_build:
	cd aws_lambda && GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap main.go && zip stori-challenge.zip bootstrap

.PHONY: test_e2e test_integration
