package service

import (
	"context"
	pb "github.com/AkezhanOb1/cf-confectioners/api/proto/confectioner"
	db "github.com/AkezhanOb1/cf-confectioners/repository"
	"github.com/golang/protobuf/ptypes/empty"
	"log"
)

//GetConfectionerById is a
func (*ConfectionerServer) GetConfectioners (ctx context.Context, empty *empty.Empty) (*pb.GetConfectionersResponse, error) {
	confectioners, err := db.GetConfectionersRepository(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}


	return confectioners, nil
}
