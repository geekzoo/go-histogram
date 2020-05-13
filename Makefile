all: version
	go build -ldflags "-X main.version=${COMMIT}"
version:
	 #$(eval COMMIT = $(shell git rev-parse HEAD))
	 $(shell echo 'go build -ldflags "-X main.version=$(git rev-parse HEAD)"')
#info:
#	$(shell echo 'go build -ldflags "-X main.version=$(git rev-parse HEAD)"')