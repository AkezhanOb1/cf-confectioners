package repository

import (
	"context"
	pb "github.com/AkezhanOb1/cf-confectioners/api/proto/confectioner"
	"github.com/AkezhanOb1/cf-confectioners/config"
	"github.com/jackc/pgx/v4"

)

//CreateConfectionerRepository is a
func CreateConfectionerRepository(ctx context.Context, request *pb.CreateConfectionerRequest) (*pb.CreateConfectionerResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)

	sqlQuery := `INSERT INTO  confectioner 
	                  (first_name, second_name, introduction, age, phone_number, 
							email, instagram_link, password, town_id)
				 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id;`


	var confectionerID int64


	row := conn.QueryRow(
		ctx,
		sqlQuery,
		request.GetFirstName(),
		request.GetSecondName(),
		request.GetIntroduction(),
		request.GetAge(),
		request.GetPhoneNumber(),
		request.GetEmail(),
		request.GetInstagramLink(),
		request.GetPassword(),
		request.GetTownId(),
	)

	err = row.Scan(&confectionerID)
	if err != nil {
		return nil, err
	}

	return &pb.CreateConfectionerResponse{
		ConfectionerId: confectionerID,
	}, nil
}