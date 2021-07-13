
GOPATH:=$(shell go env GOPATH)


.PHONY: build
build:
	GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o  im-p2p *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker: build
	docker build . -t im-p2p:latest

.PHONY: run-docker
run-docker:
	docker build . -t im-p2p:latest
	docker rm -f im-p2p | echo "remove ok"
	docker run -d --name im-p2p im-p2p
	docker ps
