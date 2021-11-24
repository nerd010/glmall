package main

import (
	"products/handler"
	pb "products/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("products"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterProductsHandler(srv.Server(), new(handler.Products))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
