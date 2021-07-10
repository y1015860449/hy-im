
GOPATH:=$(shell go env GOPATH)


.PHONY: build
build:
	GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o  im-room *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker: build
	docker build . -t im-room:latest

.PHONY: run-docker
run-docker:
	docker build . -t im-room:latest
	docker rm -f im-room | echo "remove ok"
	docker run -d --name im-room im-room
	docker ps
