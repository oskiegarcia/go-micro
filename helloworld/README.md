Just saying hello world

# Helloworld Service

Helloworld simply returns a personalized message in response to a name. Use it for testing purposes.

## How to build
```shell
make build
```


## How to run
```shell
./helloworld-srv --proxy_address=127.0.0.1:8081 --auth_id admin --auth_secret micro
```

or 
```shell
make run
```
## How to show logs
```shell
micro logs helloworld
```


## How to invoke
```shell
micro helloworld call --name=Oscar 
```
 or
 ```shell
curl "http://localhost:8080/helloworld/Call?name=Oscar"
```

## How to stop the service
```shell
micro kill helloworld
```

