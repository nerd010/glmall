package main

import (
	"coupon/handler"
	pb "coupon/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("coupon"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterCouponHandler(srv.Server(), new(handler.Coupon))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
