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

type stockController struct {
	service service.StockService
}

func NewStockController(service service.StockService) StockController {
	return &stockController{service}
}

func (c *stockController) GetData(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	days, err := strconv.ParseInt(chi.URLParam(request, "days"), 10, 32)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error parsing the date"})
		return
	}

	liters, err := c.service.GetMilkByDays(days)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the mlik stock"})
		return
	}

	skins, err := c.service.GetSkinByDays(days)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the skin stock"})
		return
	}

	var stock entity.Stock

	stock.Milk = liters
	stock.Skins = skins

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(stock)

}
