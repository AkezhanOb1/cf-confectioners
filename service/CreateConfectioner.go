package service

import (
	"context"
	"fmt"
	pb "github.com/AkezhanOb1/cf-confectioners/api/proto/confectioner"
	"github.com/AkezhanOb1/cf-confectioners/config"
	db "github.com/AkezhanOb1/cf-confectioners/repository"
	"github.com/AkezhanOb1/cf-confectioners/mq"
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

	queue, err := mq.DeclareQueue(config.ConfectionerGreetingQueue)
	if err != nil {
		return newConfectioner, fmt.Errorf("can not declare queue")
	}
	err = mq.Publish(queue, request)
	if err != nil {
		return newConfectioner, fmt.Errorf("can not send email to %s", request.GetEmail())
	}

	return newConfectioner, nil
}
