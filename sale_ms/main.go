package main

import (
	"sale_ms/handler"
	pb "sale_ms/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("sale_ms"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterSale_msHandler(srv.Server(), new(handler.Sale_ms))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
