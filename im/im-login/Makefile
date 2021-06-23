
GOPATH:=$(shell go env GOPATH)


.PHONY: build
build:
	GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o  im-login *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker: build
	docker build . -t im-login:latest

.PHONY: run-docker
run-docker:
	docker build . -t im-login:latest
	docker rm -f im-login | echo "remove ok"
	docker run -d --name im-login im-login
	docker ps
