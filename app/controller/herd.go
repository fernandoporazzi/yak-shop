package controller

import (
	"encoding/json"
	"net/http"

	"github.com/fernandoporazzi/yak-shop/app/service"
)

type HerdController interface {
	GetData(response http.ResponseWriter, request *http.Request)
}

type herdController struct {
	service service.HerdService
}

func NewHerdController(service service.HerdService) HerdController {
	return &herdController{service}
}

func (c *herdController) GetData(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)

	herd, _ := c.service.GetData(int32(13))
	json.NewEncoder(response).Encode(herd)
}
