
GOPATH:=$(shell go env GOPATH)


.PHONY: build
build:
	GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o  http-gateway *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker: build
	docker build . -t http-gateway:latest

.PHONY: run-docker
run-docker:
	docker build . -t http-gateway:latest
	docker rm -f http-gateway | echo "remove ok"
	docker run -d --name http-gateway http-gateway
	docker ps
