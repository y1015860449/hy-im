
GOPATH:=$(shell go env GOPATH)


.PHONY: build
build:
	GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o  api-friend *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker: build
	docker build . -t api-friend:latest

.PHONY: run-docker
run-docker:
	docker build . -t api-friend:latest
	docker rm -f api-friend | echo "remove ok"
	docker run -d --name api-friend api-friend
	docker ps
