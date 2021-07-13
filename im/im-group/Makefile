
GOPATH:=$(shell go env GOPATH)


.PHONY: build
build:
	GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o  im-group *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker: build
	docker build . -t im-group:latest

.PHONY: run-docker
run-docker:
	docker build . -t im-group:latest
	docker rm -f im-group | echo "remove ok"
	docker run -d --name im-group im-group
	docker ps
