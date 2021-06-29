.PHONY: run
run:
	go run --race ./cmd/main.go

.PHONY: test
test:
	go test --race -v --cover -count=1 ./...

.PHONY: test-html
test-html:
	go test --coverprofile=c.out ./...
	go tool cover --html=c.out

.DEFAULT_GOAL := run