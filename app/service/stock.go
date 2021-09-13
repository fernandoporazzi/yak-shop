package service

import (
	"math"
	"strconv"

	"github.com/fernandoporazzi/yak-shop/app/entity"
)

var (
	YakYear float64 = 100 // a Yak year consists of 100 days
)

type StockService interface {
	GetMilkByDays(days int32) (float64, error)
	GetSkinByDays(days int32) (int32, error)
}

type service struct {
	herd entity.Herd
}

func NewStockService(herd entity.Herd) StockService {
	return &service{herd}
}

func (s *service) GetMilkByDays(days int32) (float64, error) {
	var liters float64 = 0

	for _, v := range s.herd.LabYaks {
		if v.Sex != "f" {
			continue
		}

		age, err := strconv.ParseFloat(v.Age, 64)
		if err != nil {
			return 0, err
		}

		ageInYakYears := age * YakYear
		ageWithElapsedTime := ageInYakYears + float64(days)

		// Each day a LabYak produces 50 - (D * 0.03) liters of milk (D = age in days).
		l := 50 - ageWithElapsedTime*0.03

		liters = liters + l
	}

	return liters * float64(days), nil
}

func (s *service) GetSkinByDays(days int32) (int32, error) {
	var skins int32 = 0

	for _, v := range s.herd.LabYaks {
		age, err := strconv.ParseFloat(v.Age, 64)
		if err != nil {
			return 0, err
		}
		ageInYakYears := age * YakYear
		ageWithElapsedTime := ageInYakYears + float64(days)

		// At most every 8 + (D * 0.01) days you can again shave a LabYak (D = age in days).
		daysSinceLastShaved := 8 + ageWithElapsedTime*0.01

		// N.B. T=13 means that day 12 has elapsed, but day 13 has yet to begin
		s := int32(math.Ceil(float64(days-1) / daysSinceLastShaved))
		skins = skins + s
	}

	return skins, nil
}
