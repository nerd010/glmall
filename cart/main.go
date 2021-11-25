package main

import (
	"cart/handler"
	pb "cart/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("cart"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterCartHandler(srv.Server(), new(handler.Cart))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
