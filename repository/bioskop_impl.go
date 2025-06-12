package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/sb-golang-67-bioskop/database/connection"
	"github.com/sb-golang-67-bioskop/structs"
)

func NewBioskopRepository() BioskopRepositoryInterface {
	return &BioskopRepository{}
}

func (b *BioskopRepository) CreateBioskop(bioskop structs.Bioskop) (result structs.Bioskop, err error) {
	sqlStatement := `INSERT INTO bioskop (nama, lokasi, rating) VALUES ($1,$2,$3) RETURNING id`

	var id int
	errs := connection.DbConn.QueryRow(sqlStatement, bioskop.Nama, bioskop.Lokasi, bioskop.Rating).Scan(&id)
	if errs != nil {
		return structs.Bioskop{}, errs
	}

	return structs.Bioskop{
		Id:     int(id),
		Nama:   bioskop.Nama,
		Lokasi: bioskop.Lokasi,
		Rating: bioskop.Rating,
	}, nil
}

func (b *BioskopRepository) GetAllBioskop() (result []structs.Bioskop, err error) {
	var resultData []structs.Bioskop

	sqlStatement := `Select * from bioskop`

	rows, err := connection.DbConn.Query(sqlStatement)
	if err != nil {
		return resultData, err
	}

	defer rows.Close()

	for rows.Next() {
		var data = structs.Bioskop{}
		var errs = rows.Scan(&data.Id, &data.Nama, &data.Lokasi, &data.Rating)
		if err != nil {
			fmt.Println(errs.Error())
			return
		}

		resultData = append(resultData, data)
	}

	return resultData, nil
}

func (b *BioskopRepository) GetBioskopById(id string) (result structs.Bioskop, err error) {
	var resultData structs.Bioskop
	sqlStatement := `Select * from bioskop where id = $1`

	errs := connection.DbConn.QueryRow(sqlStatement, id).Scan(&resultData.Id, &resultData.Nama, &resultData.Lokasi, &resultData.Rating)
	if errs != nil {
		if errs == sql.ErrNoRows {
			var msg = "gagal menemukan user dengan id tersebut"
			return resultData, errors.New(msg)
		} else {
			return resultData, errs
		}
	}

	return resultData, nil
}

func (b *BioskopRepository) UpdateBioskop(id string, bioskop structs.Bioskop) (err error) {
	sqlStatement := `Update bioskop set nama = $2, lokasi = $3, rating = $4 where id = $1`

	res, errs := connection.DbConn.Exec(sqlStatement, id, bioskop.Nama, bioskop.Lokasi, bioskop.Rating)
	if errs != nil {
		return errs
	}

	rowsAffected, errs := res.RowsAffected()
	if errs != nil {
		return errs
	}

	if rowsAffected == 0 {
		return errors.New("data dengan id tersebut tidak ditemukan")
	}

	return nil

}

func (b *BioskopRepository) DeleteBioskop(id string) (err error) {
	sqlStatement := `Delete from bioskop where id = $1`

	res, errs := connection.DbConn.Exec(sqlStatement, id)
	if errs != nil {
		return errs
	}

	rowsAffected, errs := res.RowsAffected()
	if errs != nil {
		return errs
	}

	if rowsAffected == 0 {
		return errors.New("data dengan id tersebut tidak ditemukan")
	}

	return nil
}
