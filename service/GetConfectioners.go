package service

import (
	"context"
	pb "github.com/AkezhanOb1/cf-confectioner/api/proto/confectioner"
	db "github.com/AkezhanOb1/cf-confectioner/repository"
	empty "github.com/golang/protobuf/ptypes/empty"
	"log"
)

//GetConfectionerById is a
func (s *Server) GetConfectioners (ctx context.Context, empty *empty.Empty) (*pb.GetConfectionersResponse, error) {
	confectioners, err := db.GetConfectionersRepository(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}


	return confectioners, nil
}
