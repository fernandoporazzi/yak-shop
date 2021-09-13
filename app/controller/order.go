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

type OrderController interface {
	PlaceOrder(response http.ResponseWriter, request *http.Request)
}

type orderController struct {
	stockService service.StockService
}

func NewOrderController(stockService service.StockService) OrderController {
	return &orderController{stockService}
}

func (c *orderController) PlaceOrder(response http.ResponseWriter, request *http.Request) {
	days, err := strconv.ParseInt(chi.URLParam(request, "days"), 10, 32)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error parsing the date"})
		return
	}

	var input entity.OrderInput
	err = json.NewDecoder(request.Body).Decode(&input)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error decoding request body"})
		return
	}

	liters, err := c.stockService.GetMilkByDays(days)
	skins, err := c.stockService.GetSkinByDays(days)

	var output entity.Order
	if input.Order.Milk <= liters && input.Order.Skins <= skins {
		output.Milk = input.Order.Milk
		output.Skins = input.Order.Skins

		response.WriteHeader(http.StatusCreated)
		json.NewEncoder(response).Encode(output)
		return
	}

	if input.Order.Milk > liters && input.Order.Skins > skins {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	// Deliver only the skins because there is not enough milk in stock
	if input.Order.Milk > liters {
		response.WriteHeader(http.StatusPartialContent)

		output.Skins = input.Order.Skins

		json.NewEncoder(response).Encode(output)
	}

	// Deliver only the milk because there is not enough skins in stock
	if input.Order.Skins > skins {
		response.WriteHeader(http.StatusPartialContent)

		output.Milk = input.Order.Milk

		json.NewEncoder(response).Encode(output)
	}
}
