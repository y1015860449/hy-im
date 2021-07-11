
GOPATH:=$(shell go env GOPATH)


.PHONY: build
build:
	GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o  api-group *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker: build
	docker build . -t api-group:latest

.PHONY: run-docker
run-docker:
	docker build . -t api-group:latest
	docker rm -f api-group | echo "remove ok"
	docker run -d --name api-group api-group
	docker ps
