.PHONY: all launch download run_tests

all: run

TARGET = ./main.go

run:
	go run ${TARGET}

build:
	go build ${TARGET}

launch_race:
	go run -race ${TARGET}

download:
	go mod download

run_tests:
	go test -race ./... -cover -coverpkg ./...

check_coverage:
	go test ./... -coverprofile coverage.out
	go tool cover -html coverage.out -o coverage.html