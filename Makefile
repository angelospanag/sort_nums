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

clean:
	@rm sort_nums &> /dev/null || true
	@rm sorted_output.txt &> /dev/null || true
	@rm tmp_* &> /dev/null || true
