### NATS with Golang

# Standalone single server using Docker
To run NATS standalone server, execute the following:
```shell
docker run -p 4222:4222 -ti nats:latest
```

# Clustered
For deploying NATS cluster see following documentations.
- https://docs.nats.io/running-a-nats-service/configuration/clustering/cluster_config
- https://docs.nats.io/nats-concepts/service_infrastructure/adaptive_edge_deployment

# Embedded
To run NATS server, in Golang.
Note that it will connect to standalone server when one exists.
```go
func createServer(cfg config) {

	// options to create nats server
	opts := &natsServer.Options{
		ServerName:     cfg.ServerName,
		Host:           cfg.Host,
		Port:           cfg.Port,
		NoLog:          false,
		NoSigs:         false,
		MaxControlLine: 4096,
		MaxPayload:     65536,
	}

	// initialize the server struct
	server, err := natsServer.NewServer(opts)
	if err != nil {
		log.Fatal(err)
	}

	// run the nats server based on the options
	err = natsServer.Run(server)
	if err != nil {
		log.Fatal("Failed to start NATS server:", err)
	}

	log.Println("NATS server started")
}
```

