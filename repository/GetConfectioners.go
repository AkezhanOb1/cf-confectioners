package repository

import (
	"context"
	pb "github.com/AkezhanOb1/cf-confectioners/api/proto/confectioner"
	"github.com/AkezhanOb1/cf-confectioners/config"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"

)
//GetConfectionersRepository is a
func GetConfectionersRepository(ctx context.Context) (*pb.GetConfectionersResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)

	sqlQuery := `SELECT c.id, c.first_name, c.second_name, c.introduction, c.age, c.phone_number,
                c.email, c.instagram_link, c.town_id,
					   (SELECT ARRAY_AGG(ARRAY[ci.id::TEXT, ci.path])
				FROM confectioner_image ci WHERE ci.confectioner_id = c.id)
				FROM confectioner c;`

	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}

	var confectioners []*pb.Confectioner
	var imagesArray pgtype.TextArray

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
			&imagesArray,
		)

		if err != nil {
			return nil, err
		}

		var images []*pb.ConfectionerProfileImage
		var image pb.ConfectionerProfileImage
		for i, val := range imagesArray.Elements {
			if i % 2 == 0 {
				image.ImageId = val.String
			}else {
				func(i int, val pgtype.Text, image pb.ConfectionerProfileImage){
					image.ImagePath = val.String
					images = append(images, &image)
				}(i, val, image)
			}
		}

		confectioner.ProfilePhoto = images
		confectioners = append(confectioners, &confectioner)
	}

	return &pb.GetConfectionersResponse{
		Confectioners: confectioners,
	}, nil
}