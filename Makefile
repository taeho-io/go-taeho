# vi: ft=make

GOPATH:=$(shell go env GOPATH)

.PHONY: test
test:
	@go get github.com/rakyll/gotest
	gotest -p 1 -race -cover -v ./...

.PHONY: bench
bench:
	@go get github.com/rakyll/gotest
	gotest -p 1 -bench=.

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: generate_mocks
generate_mocks:
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen
	mockgen -package id -destination ./id/mock_id.go github.com/taeho-io/taeho-go/id ID

.PHONY: clean_mocks
clean_mocks:
	find . -name "mock_*.go" -type f -delete
	rm -rf mocks
