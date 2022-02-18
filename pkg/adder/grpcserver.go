package adder

import (
	"context"
	"github.com/Dogaev/grpc/pkg/api"
)

type GRPCServer struct {

}

func (gs *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}

//func (gs *GRPCServer) mustEmbedUnimplementedAdderServer() {
//}