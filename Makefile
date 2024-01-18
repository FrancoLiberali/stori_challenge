install_dependencies:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint:
	golangci-lint run
