package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	coupon "coupon/proto"
)

type Coupon struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Coupon) Call(ctx context.Context, req *coupon.Request, rsp *coupon.Response) error {
	log.Info("Received Coupon.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Coupon) Stream(ctx context.Context, req *coupon.StreamingRequest, stream coupon.Coupon_StreamStream) error {
	log.Infof("Received Coupon.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&coupon.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Coupon) PingPong(ctx context.Context, stream coupon.Coupon_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&coupon.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
