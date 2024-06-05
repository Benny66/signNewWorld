
init:
	go mod tidy
	go install github.com/goreleaser/goreleaser@latest
	copy .env.example .env
	
run:
	go run main.go

build:
	goreleaser release --snapshot --clean


release:
	goreleaser release --clean

.PHONY: init, run, build, release