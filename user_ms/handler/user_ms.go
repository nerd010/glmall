package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	user_ms "user_ms/proto"
)

type User_ms struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *User_ms) Call(ctx context.Context, req *user_ms.Request, rsp *user_ms.Response) error {
	log.Info("Received User_ms.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *User_ms) Stream(ctx context.Context, req *user_ms.StreamingRequest, stream user_ms.User_ms_StreamStream) error {
	log.Infof("Received User_ms.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&user_ms.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *User_ms) PingPong(ctx context.Context, stream user_ms.User_ms_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&user_ms.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
