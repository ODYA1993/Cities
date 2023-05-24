package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/DmitryOdintsov/cities/internal/store"
	"github.com/DmitryOdintsov/cities/internal/store/models"
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	Store *store.Store
}

func NewHandler(s *store.Store) *Handler {
	return &Handler{s}
}

func (h *Handler) AddCitiecHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()
	var city models.City
	if err = json.Unmarshal(content, &city); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	citySave, err := h.Store.Cities.AddCity(&city)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
	cityByte, err := json.Marshal(citySave)
	_, err = w.Write(cityByte)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	return
}

func (h *Handler) GetCitiesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cities := h.Store.Cities.GetCities()
	citiesByte, err := json.Marshal(&cities)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(citiesByte)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	return
}

func (h *Handler) GetCitiecByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(id)
	city := h.Store.Cities.GetCitiecByID(idInt)
	cityByte, err := json.Marshal(&city)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(cityByte)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	return
}

func (h *Handler) DeleteCitiecByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(id)
	city := h.Store.Cities.GetCitiecByID(idInt)
	ok := h.Store.Cities.DeleteCitiecByID(idInt)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("пользователя с таким ID нет")
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(fmt.Sprintf("город с ID %d удален", city.ID)))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	return
}

type PutPopulationInput struct {
	Population int `json:"population"`
}

func (h *Handler) PutPopulationSizeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(id)
	newPopulation, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	var cityPop *PutPopulationInput
	err = json.Unmarshal(newPopulation, &cityPop)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	fmt.Println(cityPop)
	defer r.Body.Close()
	city, err := h.Store.Cities.PutPopulationSize(idInt, cityPop.Population)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(fmt.Sprintf("Население успешно обновлено %v", city)))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	return
}

func (h *Handler) GetListByRegionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	region := chi.URLParam(r, "region")
	listRegion := h.Store.Cities.GetListByRegion(region)
	cityByte, err := json.Marshal(&listRegion)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(cityByte)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	return
}

func (h *Handler) GetListByDistrictHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	district := chi.URLParam(r, "district")
	listDistrict := h.Store.Cities.GetListByDistrict(district)
	cityByte, err := json.Marshal(&listDistrict)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(cityByte)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	return
}

func (h *Handler) GetListPopulationRangeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	from := chi.URLParam(r, "from")
	to := chi.URLParam(r, "to")
	fromInt, _ := strconv.Atoi(from)
	toInt, _ := strconv.Atoi(to)

	listPopulation := h.Store.Cities.GetListPopulationRange(fromInt, toInt)

	cityByte, err := json.Marshal(&listPopulation)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(cityByte)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	return
}

func (h *Handler) GetListCitiecYearOfFoundationRangeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	from := chi.URLParam(r, "from")
	to := chi.URLParam(r, "to")
	fromInt, _ := strconv.Atoi(from)
	toInt, _ := strconv.Atoi(to)

	listFoundation := h.Store.Cities.GetListCitiecYearOfFoundationRange(fromInt, toInt)

	cityByte, err := json.Marshal(&listFoundation)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(cityByte)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	return
}
