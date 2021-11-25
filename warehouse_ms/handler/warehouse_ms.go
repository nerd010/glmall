package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	warehouse_ms "warehouse_ms/proto"
)

type Warehouse_ms struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Warehouse_ms) Call(ctx context.Context, req *warehouse_ms.Request, rsp *warehouse_ms.Response) error {
	log.Info("Received Warehouse_ms.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Warehouse_ms) Stream(ctx context.Context, req *warehouse_ms.StreamingRequest, stream warehouse_ms.Warehouse_ms_StreamStream) error {
	log.Infof("Received Warehouse_ms.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&warehouse_ms.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Warehouse_ms) PingPong(ctx context.Context, stream warehouse_ms.Warehouse_ms_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&warehouse_ms.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
