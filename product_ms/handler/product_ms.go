package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	product_ms "product_ms/proto"
)

type Product_ms struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Product_ms) Call(ctx context.Context, req *product_ms.Request, rsp *product_ms.Response) error {
	log.Info("Received Product_ms.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Product_ms) Stream(ctx context.Context, req *product_ms.StreamingRequest, stream product_ms.Product_ms_StreamStream) error {
	log.Infof("Received Product_ms.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&product_ms.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Product_ms) PingPong(ctx context.Context, stream product_ms.Product_ms_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&product_ms.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
