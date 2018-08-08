.PHONY: format
format:
	@find . -type f -name "*.go*" -print0 | xargs -0 gofmt -s -w

.PHONY: debs
debs:
	@go get -u github.com/golang/dep/cmd/dep
	@go get -u github.com/nicksnyder/go-i18n/v2/goi18n
	@dep ensure

.PHONY: test
test:
	@go test -race ./...

.PHONY: bench
bench:
	@go test -bench=. -benchmem

.PHONY: extract
extract:
    @goi18n extract -format=json newMessages/newMessages.go

.PHONY: merge
merge:
    @goi18n merge -format=json en.json active.en.json

# Clean junk
.PHONY: clean
clean:
	@go clean ./...