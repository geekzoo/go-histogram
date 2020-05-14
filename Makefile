all: version buildbot
	go build -ldflags "-X main.version=${COMMIT}"
version:
	 $(eval COMMIT = $(shell git rev-parse HEAD))
buildbot: version
	curl -s -d "\`BUILDHOST: $(shell hostname), GIT: ${COMMIT}\`" https://bots.keybase.io/webhookbot/VJ87qB9AXhU7minAXfOzd8BoYOI

	
clean: 
	rm ${PWD}/go-histogram