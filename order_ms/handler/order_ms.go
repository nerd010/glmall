package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	order_ms "order_ms/proto"
)

type Order_ms struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Order_ms) Call(ctx context.Context, req *order_ms.Request, rsp *order_ms.Response) error {
	log.Info("Received Order_ms.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Order_ms) Stream(ctx context.Context, req *order_ms.StreamingRequest, stream order_ms.Order_ms_StreamStream) error {
	log.Infof("Received Order_ms.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&order_ms.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Order_ms) PingPong(ctx context.Context, stream order_ms.Order_ms_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&order_ms.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
