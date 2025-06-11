package bioskop

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateBioskop(ctx *gin.Context) (result Bioskop, err error)
}

type bioskopService struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &bioskopService{
		repo,
	}
}

func (s *bioskopService) CreateBioskop(ctx *gin.Context) (result Bioskop, err error) {
	var newBioskop Bioskop
	err = ctx.ShouldBind(&newBioskop)
	if err != nil {
		return
	}

	res, err := s.repo.CreateBioskopRepository(newBioskop)
	if err != nil {
		err = errors.New("gagal menambahkan data bioskop")
		return Bioskop{}, err
	}

	return res, nil
}
