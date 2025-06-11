package bioskop

import "github.com/sb-golang-67-bioskop/database/connection"

type Repository interface {
	CreateBioskopRepository(bioskop Bioskop) (result Bioskop, err error)
}

type bioskopRepository struct{}

func NewRepository() Repository{
	return &bioskopRepository{}
}

func (b *bioskopRepository) CreateBioskopRepository(bioskop Bioskop) (result Bioskop, err error) {
	sqlStatement := `INSERT INTO bioskop (nama, lokasi, rating) VALUES ($1,$2,$3) RETURNING id`

	var id int
	err = connection.DbConn.QueryRow(sqlStatement, bioskop.Nama, bioskop.Lokasi, bioskop.Rating).Scan(&id)
	if err != nil {
		return Bioskop{}, err
	}

	return Bioskop{
		Id:     int(id),
		Nama:   bioskop.Nama,
		Lokasi: bioskop.Lokasi,
		Rating: bioskop.Rating,
	}, nil
}
