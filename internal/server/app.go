package server

import (
	"github.com/DmitryOdintsov/cities/internal/server/handlers"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type App struct {
	httpServer *http.Server
}

func Run(h *handlers.Handler) {
	log.Println("starting server...")
	router := chi.NewRouter()
	router.Post("/city", h.AddCitiecHandler)
	router.Get("/cities", h.GetCitiesHandler)
	router.Get("/city/{id}", h.GetCitiecByIDHandler)
	router.Delete("/city/{id}", h.DeleteCitiecByIDHandler)
	router.Put("/city/{id}", h.PutPopulationSizeHandler)
	router.Get("/region/{region}", h.GetListByRegionHandler)
	router.Get("/district/{district}", h.GetListByDistrictHandler)
	router.Get("/population/{from}/{to}", h.GetListPopulationRangeHandler)
	router.Get("/foundation/{from}/{to}", h.GetListCitiecYearOfFoundationRangeHandler)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("starting server...")
}
