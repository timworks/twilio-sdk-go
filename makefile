TEST?=$$(go list ./...)

default: build

download: 
	@echo "==> Download dependencies"
	go mod vendor

build: fmt generate
	go install

test: fmt generate
	go test $(TESTARGS) -timeout=30s -parallel=4 $(TEST) -tags=unit

testacc: fmt generate
	go test $(TESTARGS) -timeout=30s -parallel=4 $(TEST) -tags=acceptance

fmt:
	@echo "==> Fixing source code with goimports (uses gofmt under the hood)..."
	goimports -w .

tools:
	@echo "==> installing required tooling..."
	GO111MODULE=off go get -u golang.org/x/tools/cmd/goimports
	GO111MODULE=off go get -u github.com/client9/misspell/cmd/misspell
	GO111MODULE=off go get -u github.com/gordonklaus/ineffassign
	GO111MODULE=off go get -u github.com/gojp/goreportcard/cmd/goreportcard-cli

reportcard:
	@echo "==> running go report card"
	goreportcard-cli

goreportcard-refresh:
	@echo "==> refresh goreportcard checks"
	curl -X POST -F "repo=github.com/RJPearson94/twilio-sdk-go" https://goreportcard.com/checks

generate:
	go generate  ./...

.PHONY: download build test fmt tools generate reportcard goreportcard-refresh