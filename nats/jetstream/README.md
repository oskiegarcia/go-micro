# NATS JetStream Playground

Playground for a Secure, Highly Available NATS Cluster with message persistence (using JetStream).

This repo contains a cluster with 3 nodes and a Go client sending/receiving messages to it.

## Pre-requisites
To use this repo, please install:
- [Docker](https://docs.docker.com/get-docker/)
- [mkcert](https://github.com/FiloSottile/mkcert) (zero-config tool for locally-trusted development certificates)


## Installation
Authentication for client-server and server-server (for the cluster) use X509 certificatesn. To install them locally:

```shell
make certificates
```

## Run it
Spin up the cluster with:

```shell
docker-compose up
```
Note: Use client/client.go to run on this single instance.

or

```shell
docker run -p 4222:4222 nats -js
```
Note: Use client2/client.go to run on this single instance.


Run the client with:
```shell
cd client
go run client.go
```

You should see the following appear in your terminal:
```
connecting securely to cluster
getting JetStream context
stream not found
creating stream "ORDERS" and subject "ORDERS.received"
publishing an order
attempting to receive order
got order: "one big burger"
```

You can also use `nats-io/nats-tools/nats` to issue manual commands to the cluster. If you do so, you may need to change the client publish permissions in `config/jetstream.conf`.

## Clean up
Once you are done testing, remove the CA from your local system trust store:

```shell
make cleanup
```

## Other references
- https://hackmd.io/@sofianhw/message-broker
- https://github.com/sofianhw/go-jetstream-order/blob/main/update-status/main.go

