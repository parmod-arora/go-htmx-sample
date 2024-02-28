ifeq ("$(wildcard .env)","")
    $(shell cp env.sample .env)
endif

include .env
$(eval export $(grep -v '^#' .env | xargs -0))


format:
	gofmt -s -w .
	go run ./pages fmt .


generate-template:
	templ generate

run: generate-template
	go run ./cmd/main.go

clear-exec:
	rm -rf ./main

download-deps:
	go mod download

build: download-deps generate-template
	go build ./cmd/main.go

setup-local:
	docker-compose up --build -d
