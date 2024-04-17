package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/oskiegarcia/go-micro/helloworld/handler"
	pb "github.com/oskiegarcia/go-micro/helloworld/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("helloworld"),
	)

	// Register Handler
	pb.RegisterHelloworldHandler(srv.Server(), new(handler.Helloworld))

	// Run the service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
