package service

import (
	"math"
	"strconv"

	"github.com/fernandoporazzi/yak-shop/app/entity"
)

type HerdService interface {
	GetData(days int64) (entity.HerdPayload, error)
}

type herdService struct {
	herd entity.Herd
}

func NewHerdService(herd entity.Herd) HerdService {
	return &herdService{herd}
}

func (s *herdService) GetData(days int64) (entity.HerdPayload, error) {
	var herd entity.HerdPayload
	var labyaks []entity.LabYakPayload

	for _, v := range s.herd.LabYaks {
		age, err := strconv.ParseFloat(v.Age, 64)
		if err != nil {
			return entity.HerdPayload{}, err
		}

		ageInYakYears := age * YakYear
		ageWithElapsedTime := ageInYakYears + float64(days)

		// At most every 8 + (D * 0.01) days you can again shave a LabYak (D = age in days).
		daysSinceLastShaved := 8 + ageWithElapsedTime*0.01

		var ageLastShaved float64

		// N.B. T=13 means that day 12 has elapsed, but day 13 has yet to begin
		if daysSinceLastShaved > float64(days-1) {
			ageLastShaved = ageInYakYears / YakYear
		} else {
			ageLastShaved = (ageInYakYears + math.Ceil(daysSinceLastShaved)) / YakYear
		}

		var labyak entity.LabYakPayload
		labyak.Age = ageWithElapsedTime / YakYear
		labyak.Name = v.Name
		labyak.Sex = v.Sex
		labyak.AgeLastShaved = ageLastShaved

		labyaks = append(labyaks, labyak)
	}

	herd.Herd = labyaks

	return herd, nil
}
