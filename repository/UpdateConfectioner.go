package repository

import (
	"context"
	pb "github.com/AkezhanOb1/cf-confectioner/api/proto/confectioner"
	"github.com/AkezhanOb1/cf-confectioner/config"
	"github.com/jackc/pgx/v4"

)
//UpdateConfectionerRepository is a
func UpdateConfectionerRepository(ctx context.Context, request *pb.UpdateConfectionerRequest) (*pb.UpdateConfectionerResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)
	sqlQuery := `UPDATE confectioner SET first_name=$1, second_name=$2, introduction=$3,
                        age=$4, phone_number=$5, email=$6, instagram_link=$7, town_id=$8
				WHERE id = $9;`

	_, err = conn.Exec(
		ctx,
		sqlQuery,
		request.GetFirstName(),
		request.GetSecondName(),
		request.GetIntroduction(),
		request.GetAge(),
		request.GetPhoneNumber(),
		request.GetEmail(),
		request.GetInstagramLink(),
		request.GetTownId(),
		request.GetId(),
		)

	if err != nil {
		return &pb.UpdateConfectionerResponse{
			Updated: false,
		}, err
	}

	return &pb.UpdateConfectionerResponse{
		Updated: true,
	}, nil
}