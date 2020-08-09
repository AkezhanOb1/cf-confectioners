package service

import (
	"context"
	pb "github.com/AkezhanOb1/cf-confectioners/api/proto/confectioner"
	db "github.com/AkezhanOb1/cf-confectioners/repository"
	"github.com/AkezhanOb1/cf-confectioners/service/helper"
	"log"
)

//GetConfectionerById is a
func (*ConfectionerServer) CreateConfectioner (ctx context.Context, request *pb.CreateConfectionerRequest) (*pb.CreateConfectionerResponse, error) {

	hashedPassword, err := helper.HashPassword(request.Password)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	request.Password = hashedPassword

	newConfectioner, err := db.CreateConfectionerRepository(ctx, request)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return newConfectioner, nil
}
