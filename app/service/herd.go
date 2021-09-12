package service

import (
	"github.com/fernandoporazzi/yak-shop/app/entity"
)

type HerdService interface {
	GetData(days int32) (entity.Herd, error)
}

type herdService struct {
	herd entity.Herd
}

func NewHerdService(herd entity.Herd) HerdService {
	return &herdService{herd}
}

func (s *herdService) GetData(days int32) (entity.Herd, error) {
	return s.herd, nil
}
