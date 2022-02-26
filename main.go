package main

import (
	"github.com/myself-demo/order-service-api/pb"
	"github.com/myself-demo/order-service/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	pb.RegisterOrderServiceServer(server, service.NewOrderServiceProvider())
	if err := server.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
