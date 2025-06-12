package service

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/sb-golang-67-bioskop/repository"
	"github.com/sb-golang-67-bioskop/structs"
)

func NewBioskopService(repo repository.BioskopRepositoryInterface) BioskopServiceInterface {
	return &BioskopService{
		repo,
	}
}

func (s *BioskopService) CreateBioskop(ctx *gin.Context) (result structs.Bioskop, err error) {
	var newBioskop structs.Bioskop
	err = ctx.ShouldBind(&newBioskop)
	if err != nil {
		if len(ctx.Errors) > 0 {
			return structs.Bioskop{}, errors.New(ctx.Errors.Last().Error())
		}
		return
	}

	res, err := s.repo.CreateBioskop(newBioskop)
	if err != nil {
		return structs.Bioskop{}, err
	}

	return res, nil
}

func (s *BioskopService) GetAllBioskop(ctx *gin.Context) (result []structs.Bioskop, err error) {
	res, err := s.repo.GetAllBioskop()
	if err != nil {
		return []structs.Bioskop{}, err
	}

	return res, nil
}

func (s *BioskopService) GetBioskopById(ctx *gin.Context) (result structs.Bioskop, err error) {
	id := ctx.Param("id")

	res, err := s.repo.GetBioskopById(id)
	if err != nil {
		return structs.Bioskop{}, err
	}

	return res, nil
}

func (s *BioskopService) UpdateBioskop(ctx *gin.Context) (err error) {
	var updateBioskop structs.Bioskop
	errs := ctx.ShouldBindJSON(&updateBioskop)
	if errs != nil {
		return errs
	}

	var id = ctx.Param("id")
	errs = s.repo.UpdateBioskop(id, updateBioskop)
	if errs != nil {
		return errs
	}

	return nil
}

func (s *BioskopService) DeleteBioskop(ctx *gin.Context) (err error) {
	id := ctx.Param("id")

	errs := s.repo.DeleteBioskop(id)
	if errs != nil {
		return errs
	}

	return nil
}
