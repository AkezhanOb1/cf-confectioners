package service

import (
	"context"
	pb "github.com/AkezhanOb1/cf-confectioner/api/proto/confectioner"
	empty "github.com/golang/protobuf/ptypes/empty"
)

//GetConfectionerById is a
func (s *Server) GetConfectioners (ctx context.Context, empty *empty.Empty) (*pb.GetConfectionersResponse, error) {
	return nil, nil
}
