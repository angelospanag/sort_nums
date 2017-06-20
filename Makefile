.PHONY: build doc imports run test
IMPORTSTARGET=$$(find . -type f -name '*.go' -not -path "./vendor/*")
TESTTARGET=$$(go list ./... | grep -v /vendor/)
APPBIN=$(shell basename $(PWD))

imports:
	@goimports -w $(IMPORTSTARGET)

build: imports
	@go build

doc:
	@godoc -http=:8080 -index

run: build
	@./$(APPBIN)

test: imports
	@go test -timeout=5s -cover -race $(TESTTARGET)
