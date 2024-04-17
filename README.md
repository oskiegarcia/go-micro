# Guide


## Install proto dependencies
```shell
go install github.com/golang/protobuf/protoc-gen-go@latest
go install github.com/micro/micro/v3/cmd/protoc-gen-micro@latest
```

## Run the server
```shell
micro server
```

## How to list services
```shell
micro services
```

### micro dashboard
```shell
micro web
```
and then open http://localhost:8082/


## Examples
### Helloworld
https://github.com/micro/services/tree/master/helloworld


