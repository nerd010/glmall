package main

import (
	"item/handler"
	pb "item/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("item"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterItemHandler(srv.Server(), new(handler.Item))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
