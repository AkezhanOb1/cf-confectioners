package service

import (
	"context"
	pb "github.com/AkezhanOb1/cf-confectioners/api/proto/confectioner"
	db "github.com/AkezhanOb1/cf-confectioners/repository"
	"log"
)

//GetConfectionerById is a
func (*ConfectionerServer) UpdateConfectioner (ctx context.Context, request *pb.UpdateConfectionerRequest) (*pb.UpdateConfectionerResponse, error) {
	updated, err := db.UpdateConfectionerRepository(ctx, request)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return updated, nil
}
