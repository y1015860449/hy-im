
GOPATH:=$(shell go env GOPATH)


.PHONY: build
build:
	GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o  im-gateway *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker: build
	docker build . -t im-gateway:latest

.PHONY: run-docker
run-docker:
	docker build . -t im-gateway:latest
	docker rm -f im-gateway | echo "remove ok"
	docker run -d --name im-gateway im-gateway
	docker ps
