all: version
	go build -ldflags "-X main.version=${COMMIT}"
version:
	 $(eval COMMIT = $(shell git rev-parse HEAD))