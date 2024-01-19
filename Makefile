install_dependencies:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/vektra/mockery/v2@v2.40.1

lint:
	golangci-lint run
	cd test_e2e && golangci-lint run --config ../.golangci.yml

test_unit:
	go test -v ./...

test_e2e:
	go install .
	go test ./test_e2e

.PHONY: test_e2e
