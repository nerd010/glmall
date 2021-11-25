package main

import (
	"user_ms/handler"
	pb "user_ms/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("user_ms"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterUser_msHandler(srv.Server(), new(handler.User_ms))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
