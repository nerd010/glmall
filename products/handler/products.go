package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	products "products/proto"
)

type Products struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Products) Call(ctx context.Context, req *products.Request, rsp *products.Response) error {
	log.Info("Received Products.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Products) Stream(ctx context.Context, req *products.StreamingRequest, stream products.Products_StreamStream) error {
	log.Infof("Received Products.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&products.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Products) PingPong(ctx context.Context, stream products.Products_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&products.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
