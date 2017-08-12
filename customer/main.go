package main

import (
	"log"
	"net"

	"os"
	"os/signal"
	"syscall"

	"github.com/vasiliy-t/microtrain/customer/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Get(ctx context.Context, req *proto.GetCustomerRequest) (*proto.GetCustomerResponse, error) {
	return &proto.GetCustomerResponse{}, nil
}

func (s *server) Save(ctx context.Context, req *proto.SaveCustomerRequest) (*proto.SaveCustomerResponse, error) {
	return &proto.SaveCustomerResponse{}, nil
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Fatalf("failed to listen %s", err)
	}

	s := grpc.NewServer()
	proto.RegisterCustomerServiceServer(s, &server{})
	log.Println("customer service starting")

	exitchan := make(chan os.Signal)
	signal.Notify(exitchan, syscall.SIGTERM)

	go func() {
		err = s.Serve(l)
		if err != nil {
			log.Fatalf("Error while starting grpc server %s", err)
		}
	}()
	<-exitchan
	s.GracefulStop()
}
