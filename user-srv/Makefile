GOPATH:=$(shell go env GOPATH)
.PHONY:
run-loc:
	go run main.go -env=loc

.PHONY: build
build: proto
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../build/user-srv main.go plugin.go

.PHONY: docker
docker:docker
	docker build . -t user-srv:latest