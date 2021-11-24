package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	ware "ware/proto"
)

type Ware struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Ware) Call(ctx context.Context, req *ware.Request, rsp *ware.Response) error {
	log.Info("Received Ware.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Ware) Stream(ctx context.Context, req *ware.StreamingRequest, stream ware.Ware_StreamStream) error {
	log.Infof("Received Ware.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&ware.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Ware) PingPong(ctx context.Context, stream ware.Ware_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&ware.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
