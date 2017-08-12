package main

import (
	"log"
	"net"

	"github.com/vasiliy-t/microtrain/wishlist/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Get(ctx context.Context, req *proto.GetRequest) (*proto.GetResponse, error) {
	return &proto.GetResponse{}, nil
}

func main() {
	log.Println("starting wishlist service")
	l, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Fatalf("Failed to listen %s", err)
	}

	s := grpc.NewServer()
	proto.RegisterWishlistServiceServer(s, &server{})

	err = s.Serve(l)
	if err != nil {
		log.Fatalf("Failed to start grpc server %s", err)
	}
}
