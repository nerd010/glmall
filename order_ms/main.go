package main

import (
	"order_ms/handler"
	pb "order_ms/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("order_ms"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterOrder_msHandler(srv.Server(), new(handler.Order_ms))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
