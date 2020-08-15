package main

import (
	pb "github.com/AkezhanOb1/cf-confectioners/api/proto/confectioner"
	"github.com/AkezhanOb1/cf-confectioners/config"
	"github.com/AkezhanOb1/cf-confectioners/mq"
	"github.com/AkezhanOb1/cf-confectioners/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	defer mq.Channel.Close()

	lis, err := net.Listen("tcp", config.ServerAddress)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	log.Printf("Server is listening on %v ...", config.ServerAddress)

	s := grpc.NewServer()
	pb.RegisterConfectionerServiceServer(s, &service.ConfectionerServer{})
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}