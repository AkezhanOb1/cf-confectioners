package repository

import (
	"context"
	pb "github.com/AkezhanOb1/cf-confectioner/api/proto/confectioner"
	"github.com/AkezhanOb1/cf-confectioner/config"
	"github.com/jackc/pgx/v4"

)
//GetConfectionersRepository is a
func GetConfectionersRepository(ctx context.Context) (*pb.GetConfectionersResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)

	sqlQuery := `SELECT id, first_name, second_name, introduction, age, phone_number, 
	email, instagram_link, town_id FROM confectioner;`

	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}

	var confectioners []*pb.Confectioner

	for rows.Next() {
		var confectioner pb.Confectioner

		err = rows.Scan(
			&confectioner.Id,
			&confectioner.FirstName,
			&confectioner.SecondName,
			&confectioner.Introduction,
			&confectioner.Age,
			&confectioner.PhoneNumber,
			&confectioner.Email,
			&confectioner.InstagramLink,
			&confectioner.TownId,
		)

		if err != nil {
			return nil, err
		}

		confectioners = append(confectioners, &confectioner)
	}



	return &pb.GetConfectionersResponse{
		Confectioners: confectioners,
	}, nil
}