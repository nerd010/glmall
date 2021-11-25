package main

import (
	"warehouse_ms/handler"
	pb "warehouse_ms/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("warehouse_ms"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterWarehouse_msHandler(srv.Server(), new(handler.Warehouse_ms))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
