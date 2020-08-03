package service

import (
"context"
pb "github.com/AkezhanOb1/cf-confectioner/api/proto/confectioner"
db "github.com/AkezhanOb1/cf-confectioner/repository"
	"log"
)


//GetConfectionerById is a
func (s *Server) GetConfectionerById (ctx context.Context, request *pb.GetConfectionerByIdRequest) (*pb.GetConfectionerByIdResponse, error) {
	confectioner, err := db.GetConfectionerByIdRepository(ctx, request)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return confectioner, nil
}