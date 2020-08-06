package repository

import (
	"context"
	pb "github.com/AkezhanOb1/cf-confectioner/api/proto/confectioner"
	"github.com/AkezhanOb1/cf-confectioner/config"
	"github.com/jackc/pgx/v4"

)
//GetConfectionerByIdRepository is a
func DeleteConfectionerRepository(ctx context.Context, request *pb.DeleteConfectionerRequest) (*pb.DeleteConfectionerResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)
	sqlQuery := `DELETE FROM confectioner WHERE id=$1;`

	_, err = conn.Exec(ctx, sqlQuery, request.GetConfectionerId())

	if err != nil {
		return &pb.DeleteConfectionerResponse{
			Deleted: false,
		}, err
	}

	return &pb.DeleteConfectionerResponse{
		Deleted: true,
	}, nil
}