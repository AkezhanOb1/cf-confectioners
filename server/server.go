package main

import (
	pb "github.com/AkezhanOb1/cf-confectioner/api/proto/confectioner"
	"github.com/AkezhanOb1/cf-confectioner/config"
	"google.golang.org/grpc"
	"log"
	"net"

	"github.com/AkezhanOb1/cf-confectioner/service"
)

func main() {
	lis, err := net.Listen("tcp", config.ServerAddress)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	log.Printf("Server is listening on %v ...", config.ServerAddress)

	s := grpc.NewServer()
	pb.RegisterConfectionerServiceServer(s, &service.Server{})
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}