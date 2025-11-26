default: # list recepies
    just --list

# Clean working tree
clean:
	find . -name '*_mock.go' -delete
	find . -name '*_enum.go' -delete
	find . -name '*_gen.go' -delete

# Prepare build prerequisities
prepare:
	go generate ./...
	go get ./...

# Run tests
test:
	GOEXPERIMENT=jsonv2 go test -race -cover -coverprofile=lapsus.coverage ./...

# Show code coverage report
coverage:
	GOEXPERIMENT=jsonv2 go tool cover -html=lapsus.coverage

# Executes golangci linter
lint: fmt
	golangci-lint run --path-prefix $(pwd)
	git diff --exit-code

# Format source code
fmt:
	go fmt ./...
	goimports-reviser -imports-order std,general,company,project -company-prefixes=github.com/aerario ./...
