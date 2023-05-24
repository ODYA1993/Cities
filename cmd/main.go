package main

import (
	"fmt"
	"github.com/DmitryOdintsov/cities/internal/server"
	"github.com/DmitryOdintsov/cities/internal/server/handlers"
	"github.com/DmitryOdintsov/cities/internal/store"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

var (
	path = "cities.csv"
)

func main() {
	cities := store.NewCities()
	newStore := store.NewStore(cities)
	handler := handlers.NewHandler(newStore)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	err = cities.ParseCsv(file)
	if err != nil {
		log.Println(err)
	}

	go func() {
		server.Run(handler)
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-done:
				newStore.Cities.Writer(path)
				fmt.Println("exit")
				return
			default:
				time.Sleep(time.Second * 2)
			}
		}
	}()
	wg.Wait()
}
