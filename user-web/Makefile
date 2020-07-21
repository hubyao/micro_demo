  
GOPATH:=$(shell go env GOPATH)

.PHONY: 
build: 
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../build/user-web/user-web main.go
	cp -r ./conf ../build/user-web
	cp Dockerfile  ../build/user-web

.PHONY: docker
docker:docker
	docker build ../build/user-web -t user-web:latest

## 运行docker
## 本地运行