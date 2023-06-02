.PHONY: build
build:
	go get github.com/google/wire/internal/wire
	go get github.com/google/wire/cmd/wire
	cd main; go run github.com/google/wire/cmd/wire
	go mod tidy
	sam build
