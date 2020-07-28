 GOPATH:=$(shell go env GOPATH)
.PHONY:
run-loc:
	go run main.go -env=loc

.PHONY:
build: 
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../build/user-web main.go

.PHONY: docker
docker:docker
	docker build . -t user-web:latest