package app

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/fernandoporazzi/yak-shop/app/controller"
	"github.com/fernandoporazzi/yak-shop/app/entity"
	"github.com/fernandoporazzi/yak-shop/app/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Start() {
	xmlFile, err := os.Open("herd.xml")
	if err != nil {
		fmt.Println("Error")
		panic(err)
	}
	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error creating byteValue")
		panic(err)
	}

	var herd entity.Herd

	err = xml.Unmarshal(byteValue, &herd)
	if err != nil {
		fmt.Println("Error unmarshalling xml")
		panic(err)
	}

	stockService := service.NewStockService(herd)
	herdService := service.NewHerdService(herd)

	stockController := controller.NewStockController(stockService)
	herdController := controller.NewHerdController(herdService)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("The YakShop API"))
	})

	r.Route("/yak-shop", func(r chi.Router) {
		r.Get("/stock/{days:[0-9]+}", stockController.GetData)
		r.Get("/herd/{days:[0-9]+}", herdController.GetData)
	})

	http.ListenAndServe(":3000", r)
}
