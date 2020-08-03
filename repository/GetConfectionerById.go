package repository

import (
	"context"
	pb "github.com/AkezhanOb1/cf-confectioner/api/proto/confectioner"
	"github.com/AkezhanOb1/cf-confectioner/config"
	"github.com/jackc/pgx/v4"

)
//GetConfectionerByIdRepository is a
func GetConfectionerByIdRepository(ctx context.Context, request *pb.GetConfectionerByIdRequest) (*pb.GetConfectionerByIdResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)
	sqlQuery := `SELECT id, first_name, second_name, introduction, age, phone_number, 
	email, instagram_link FROM confectioner WHERE id=$1;`

	row := conn.QueryRow(ctx, sqlQuery, request.GetId())

	var confectioner pb.Confectioner

	err = row.Scan(
		&confectioner.Id,
		&confectioner.FirstName,
		&confectioner.SecondName,
		&confectioner.Introduction,
		&confectioner.Age,
		&confectioner.PhoneNumber,
		&confectioner.Email,
		&confectioner.InstagramLink,
	)

	if err != nil {
		return nil, err
	}

	return &pb.GetConfectionerByIdResponse{
		Confectioner: &confectioner,
	}, nil
}