all: version buildbot docker_build test
	#go build -ldflags "-X main.version=${COMMIT}"
version:
	 $(eval COMMIT = $(shell git rev-parse HEAD))
buildbot: version
	curl -s -d "\`BUILDHOST: $(shell hostname), GIT: ${COMMIT}\`" https://bots.keybase.io/webhookbot/VJ87qB9AXhU7minAXfOzd8BoYOI
dev:
	go build #-ldflags "-X main.version=${COMMIT}"
docker_build: version
	go build -ldflags "-X main.version=${COMMIT}"
	docker build -t dc.geekzoo.io/go-histogram:${COMMIT} .
docker_run: version docker_build
	docker run -it --rm dc.geekzoo.io/go-histogram:${COMMIT}
test: version
	go run histogram.go -v
	printf "abc 123 \n def 32 \n" | go run histogram.go
clean: version
	rm ${PWD}/go-histogram
	docker rmi dc.geekzoo.io/go-histogram:${COMMIT}