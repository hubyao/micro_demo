.PHONY: run
run:
	go run main.go --registry=etcd --registry_address=127.0.0.1:2379 api -handler=http  --namespace=mu.micro.book.api

.PHONY:
build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o  ../build/gateway main.go

.PHONY: docker
docker:docker
	docker build . -t gateway:latest