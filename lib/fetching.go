package lib

import (
	"encoding/json"
	"io"
	"net/http"
)

// Fetch and parse the API data
func fetchData() (*Data, error) {
	resp, err := http.Get("http://localhost:3000/api")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiEndpoints map[string]string
	err = json.Unmarshal(body, &apiEndpoints)
	if err != nil {
		return nil, err
	}

	type result[T any] struct {
		data T
		err error
	}
	
	// Fetch all data with go routines
	manuChan := make(chan result[[]Manufacturer])
	catChan := make(chan result[[]Category])
	modelChan := make(chan result[[]CarModel])

	go func() {
		data, err := fetchManufacturers(apiEndpoints["manufacturers"])
		manuChan <- result[[]Manufacturer]{data, err}
	}()

	go func() {
		data, err := fetchCategories(apiEndpoints["categories"])
		catChan <- result[[]Category]{data, err}
	}()

	go func() {
		data, err := fetchCarModels(apiEndpoints["models"])
		modelChan <- result[[]CarModel]{data, err}
	}()

	manufacturersRes := <-manuChan
	categoriesRes := <-catChan
	carModelsRes := <-modelChan

	if manufacturersRes.err != nil {
		return nil, manufacturersRes.err
	}

	if categoriesRes.err != nil {
		return nil, categoriesRes.err
	}

	if carModelsRes.err != nil {
		return nil, carModelsRes.err
	}

	return &Data{
		Manufacturers: manufacturersRes.data,
		Categories: categoriesRes.data,
		CarModels: carModelsRes.data,
	}, nil
}

func fetchManufacturers(url string) ([]Manufacturer, error) {
	resp, err := http.Get("http://localhost:3000" + url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var manufacturers []Manufacturer
	json.Unmarshal(body, &manufacturers)
	return manufacturers, nil
}

func fetchCategories(url string) ([]Category, error) {
	resp, err := http.Get("http://localhost:3000" + url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var categories []Category
	json.Unmarshal(body, &categories)
	return categories, nil
}

func fetchCarModels(url string) ([]CarModel, error) {
	resp, err := http.Get("http://localhost:3000" + url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var carModels []CarModel
	json.Unmarshal(body, &carModels)
	return carModels, nil
}