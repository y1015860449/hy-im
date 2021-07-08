
GOPATH:=$(shell go env GOPATH)


.PHONY: build
build:
	GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o  im-login *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker: build
	docker build . -t api-user:latest

.PHONY: run-docker
run-docker:
	docker build . -t api-user:latest
	docker rm -f api-user | echo "remove ok"
	docker run -d --name api-user api-user
	docker ps
