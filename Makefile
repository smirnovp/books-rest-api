.PHONY: run
run:
	go run --race ./cmd/main.go

.PHONY: test
test:
	go test --race --cover -count=1 ./... | sed ''/PASS/s//$$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/''

.PHONY: test-v
test-v:
	go test --race --cover -v -count=1 ./... | sed ''/PASS/s//$$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/''

.PHONY: test-html
test-html:
	go test --coverprofile=c.out ./...
	go tool cover --html=c.out

.PHONY: run-docker-db
run-docker-db:
	docker run --rm -d -e POSTRGRES_USER=postgres -e POSTGRES_PASSWORD=password -p 5432:5432 --name local-postgres postgres


.DEFAULT_GOAL := run