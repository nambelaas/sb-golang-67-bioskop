package repository

import "github.com/sb-golang-67-bioskop/structs"

type BioskopRepositoryInterface interface {
	CreateBioskop(bioskop structs.Bioskop) (result structs.Bioskop, err error)
	GetAllBioskop() (result []structs.Bioskop, err error)
	GetBioskopById(id string) (result structs.Bioskop, err error)
	UpdateBioskop(id string, bioskop structs.Bioskop) (err error)
	DeleteBioskop(id string) (err error)
}

type BioskopRepository struct{}
