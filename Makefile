install_dependencies:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint:
	golangci-lint run
	cd test_e2e && golangci-lint run --config ../.golangci.yml

test_unit:
	go test -v ./...
