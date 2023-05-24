package store

import (
	"encoding/csv"
	"errors"
	"github.com/DmitryOdintsov/cities/internal/store/models"
	"log"
	"os"
	"strconv"
	"sync"
)

type Cities struct {
	Cities []*models.City
	sync.Mutex
}

func NewCities() *Cities {
	return &Cities{}
}

func (c *Cities) AddCity(city *models.City) (*models.City, error) {
	s := c.Cities
	s = append(s, city)
	c.Cities = s
	return city, nil
}

func (c *Cities) GetCities() []*models.City {
	var result []*models.City
	for _, v := range c.Cities {
		result = append(result, v)

	}
	return result
}

func (c *Cities) GetCitiecByID(id int) *models.City {
	for _, v := range c.Cities {
		if v.ID == id {
			return v
		}
	}
	return nil
}

func (c *Cities) PutPopulationSize(id, population int) (*models.City, error) {
	for _, v := range c.Cities {
		if v.ID == id {
			v.Population = population
			return v, nil
		}

	}
	return nil, errors.New("город не найден")
}

func (c *Cities) GetListByRegion(region string) []*models.City {
	var result []*models.City
	for _, v := range c.Cities {
		if v.Region == region {
			result = append(result, v)
		}
	}

	return result
}

func (c *Cities) GetListByDistrict(district string) []*models.City {
	var result []*models.City
	for _, v := range c.Cities {
		if v.District == district {
			result = append(result, v)
		}
	}
	return result
}

func (c *Cities) GetListPopulationRange(from, to int) []*models.City {
	var result []*models.City
	for _, v := range c.Cities {
		if v.Population >= from && v.Population <= to {
			result = append(result, v)
		}
	}
	return result
}

func (c *Cities) GetListCitiecYearOfFoundationRange(from int, to int) []*models.City {
	var result []*models.City
	for _, v := range c.Cities {
		if v.Foundation >= from && v.Foundation <= to {
			result = append(result, v)
		}
	}
	return result
}

func (c *Cities) DeleteCitiecByID(id int) bool {
	s := c.Cities
	for i, v := range c.Cities {
		if v.ID == id {
			s = append(s[:i], s[i+1:]...)
			c.Cities = s
		}
	}
	return true
}

func (c *Cities) CreateCitiecList(data [][]string) *Cities {
	var cities []*models.City
	for i, line := range data {
		if i >= 0 {
			city := models.City{}
			for j, field := range line {
				switch j {
				case 0:
					fieldInt, _ := strconv.Atoi(field)
					city.ID = fieldInt
				case 1:
					city.Name = field
				case 2:
					city.Region = field
				case 3:
					city.District = field
				case 4:
					fieldInt, _ := strconv.Atoi(field)
					city.Population = fieldInt
				case 5:
					fieldInt, _ := strconv.Atoi(field)
					city.Foundation = fieldInt
				}
			}
			cities = append(cities, &city)
		}
	}
	c.Cities = cities
	return c
}

func (c *Cities) ParseCsv(file *os.File) error {
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	c.CreateCitiecList(data)
	return nil
}

func (c *Cities) Writer(path string) {
	cities := c.GetCities()
	sliceCities := c.CitiesString(cities)

	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	write := csv.NewWriter(file)

	defer write.Flush()

	write.WriteAll(sliceCities)
}

func (c *Cities) CitiesString(cities []*models.City) [][]string {
	var result [][]string
	for _, v := range cities {
		sl := []string{
			strconv.Itoa(v.ID),
			v.Name,
			v.Region,
			v.District,
			strconv.Itoa(v.Population),
			strconv.Itoa(v.Foundation),
		}
		result = append(result, sl)
	}
	return result
}
