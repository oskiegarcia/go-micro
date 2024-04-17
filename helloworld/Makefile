
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --openapi_out=. --proto_path=.:. --micro_out=. --go_out=. proto/helloworld.proto

.PHONY: api
api:
	protoc --openapi_out=. --proto_path=.:. proto/helloworld.proto

.PHONY: build
build: proto

	go build -o helloworld-srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t helloworld-srv:latest


.PHONY: run
run:
	./helloworld-srv --proxy_address=127.0.0.1:8081 --auth_id admin --auth_secret micro