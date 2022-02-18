package main

import (
	"github.com/Dogaev/grpc/pkg/adder"
	"github.com/Dogaev/grpc/pkg/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}

	api.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen server on port 8080: %v", err)
	}

	if err = s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}