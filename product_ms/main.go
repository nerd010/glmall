package main

import (
	"product_ms/handler"
	pb "product_ms/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("product_ms"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterProduct_msHandler(srv.Server(), new(handler.Product_ms))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
