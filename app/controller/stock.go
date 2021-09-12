package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fernandoporazzi/yak-shop/app/entity"
	"github.com/fernandoporazzi/yak-shop/app/errors"
	"github.com/fernandoporazzi/yak-shop/app/service"
	"github.com/go-chi/chi/v5"
)

type StockController interface {
	GetData(response http.ResponseWriter, request *http.Request)
}

type controller struct {
	service service.StockService
}

func NewStockController(service service.StockService) StockController {
	return &controller{service}
}

func (c *controller) GetData(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	days, err := strconv.ParseInt(chi.URLParam(request, "days"), 10, 32)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error parsing the date"})
	}

	liters, err := c.service.GetMilkByDays(int32(days))

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the mlik stock"})
	}

	skins, err := c.service.GetSkinByDays(int32(days))
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the skin stock"})
	}

	var stock entity.Stock

	stock.Milk = int32(liters)
	stock.Skins = skins

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(stock)

}
