package service

import (
	"context"
	pb "github.com/AkezhanOb1/cf-confectioner/api/proto/confectioner"
	db "github.com/AkezhanOb1/cf-confectioner/repository"
	"log"
)

//GetConfectionerById is a
func (s *Server) DeleteConfectioner (ctx context.Context, request *pb.DeleteConfectionerRequest) (*pb.DeleteConfectionerResponse, error) {
	deleted, err := db.DeleteConfectionerRepository(ctx, request)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return deleted, nil
}
