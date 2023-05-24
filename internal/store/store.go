package store

import (
	"github.com/DmitryOdintsov/cities/internal/store/models"
	"os"
)

type ICities interface {
	AddCity(city *models.City) (*models.City, error)
	GetCities() []*models.City
	GetCitiecByID(id int) *models.City
	DeleteCitiecByID(id int) bool
	PutPopulationSize(id int, population int) (*models.City, error)
	GetListByRegion(region string) []*models.City
	GetListByDistrict(region string) []*models.City
	GetListPopulationRange(from int, to int) []*models.City
	GetListCitiecYearOfFoundationRange(from int, to int) []*models.City
	CreateCitiecList(data [][]string) *Cities
	ParseCsv(file *os.File) error
	Writer(path string)
	CitiesString(cities []*models.City) [][]string
}

type Store struct {
	Cities ICities
}

func NewStore(cities ICities) *Store {
	return &Store{cities}
}
