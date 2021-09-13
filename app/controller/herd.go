package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fernandoporazzi/yak-shop/app/errors"
	"github.com/fernandoporazzi/yak-shop/app/service"
	"github.com/go-chi/chi/v5"
)

type HerdController interface {
	GetData(response http.ResponseWriter, request *http.Request)
}

type herdController struct {
	herdService service.HerdService
}

func NewHerdController(service service.HerdService) HerdController {
	return &herdController{service}
}

func (c *herdController) GetData(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)

	days, err := strconv.ParseInt(chi.URLParam(request, "days"), 10, 32)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error parsing the date"})
		return
	}

	herd, _ := c.herdService.GetData(days)
	json.NewEncoder(response).Encode(herd)
}
