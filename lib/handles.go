package lib

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(render500Page()))
		}
	}()

	allData, err := fetchData()
	if err != nil {
		panic("Error while fetching data!")
	}

	queryText := r.URL.Query().Get("search")

	// Read query parameters
	selectedFilters := map[string]string{
		"search": 		queryText,
		"manufacturer": r.URL.Query().Get("manufacturer"),
		"category":     r.URL.Query().Get("category"),
		"year":         r.URL.Query().Get("year"),
		"horsepower":   r.URL.Query().Get("horsepower"),
		"engine":       r.URL.Query().Get("engine"),
	}

	manufacturerID, _ := strconv.Atoi(selectedFilters["manufacturer"])
	categoryID, _ := strconv.Atoi(selectedFilters["category"])
	year, _ := strconv.Atoi(selectedFilters["year"])
	engine := selectedFilters["engine"]

	var minHP, maxHP int
	if selectedFilters["horsepower"] != "" {
		_, err := fmt.Sscanf(selectedFilters["horsepower"], "%d-%d", &minHP, &maxHP)
		if err != nil {
			minHP = 0
			maxHP = 10000
		}
	}

	// Filter the car models
	var filteredCars []CarModel
	for _, car := range allData.CarModels {
		if (manufacturerID == 0 || car.ManufacturerID == manufacturerID) &&
			(categoryID == 0 || car.CategoryID == categoryID) &&
			(year == 0 || car.Year == year) &&
			(selectedFilters["horsepower"] == "" || 
			(car.Specifications.Horsepower >= minHP && car.Specifications.Horsepower <= maxHP)) &&
			(engine == "" || car.Specifications.Engine == engine) &&
			(queryText == "" || strings.Contains(strings.ToLower(car.Name), strings.ToLower(queryText))) {
			filteredCars = append(filteredCars, car)
		}
	}

	// Create a filtered data object
	filteredData := &Data{
		Manufacturers: allData.Manufacturers,
		Categories:    allData.Categories,
		CarModels:     filteredCars,
	}

	recentlyViewedCars := getRecentlyViewedCars(r, allData)

	// Render the page with filtered data
	fmt.Fprint(w, renderCarList(allData, filteredData, recentlyViewedCars, selectedFilters))
}

func CarHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(render500Page()))
		}
	}()

	query := r.URL.Query()
	idStr := query.Get("id")

	if idStr == "" {
		http.Error(w, "Car ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}

	data, err := fetchData()
	if err != nil {
		panic("Error while fetching data!")
	}

	// Find the requested car
	var selectedCar *CarModel
	for _, car := range data.CarModels {
		if car.ID == id {
			selectedCar = &car
			break
		}
	}

	if selectedCar == nil {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}

	addToRecentlyViewed(w, r, id)

	// Render and display the car details page
	html := renderCarDetails(*selectedCar, data.Manufacturers, data.Categories)
	fmt.Fprint(w, html)
}

func ComparePageHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(render500Page()))
		}
	}()

	// Simulated HTTP500 error -> http://localhost:8080/compare?car1=9999
	if r.URL.Query().Get("car1") == "9999" {
		panic("simulated error")
	}

	data, err := fetchData()
	if err != nil {
		panic("Error while fetching data!")
	}

	// Get selected car IDs from query parameters
	id1, _ := strconv.Atoi(r.URL.Query().Get("car1"))
	id2, _ := strconv.Atoi(r.URL.Query().Get("car2"))

	var car1, car2 *CarModel
	for _, car := range data.CarModels {
		if car.ID == id1 {
			car1 = &car
		}
		if car.ID == id2 {
			car2 = &car
		}
	}

	html := renderComparisonPage(data.CarModels, data.Manufacturers, data.Categories, car1, car2)
	fmt.Fprint(w, html)
}

func addToRecentlyViewed(w http.ResponseWriter, r *http.Request, carID int) {
    const cookieName = "recently_viewed"
    const maxItems = 5

    var ids []string
    idStr := strconv.Itoa(carID)

    // Get current cookie
    cookie, err := r.Cookie(cookieName)
    if err == nil {
        ids = strings.Split(cookie.Value, ",")
    }

    // Remove existing instance if it exists
    var newList []string
    for _, id := range ids {
        if id != idStr {
            newList = append(newList, id)
        }
    }

    // Prepend the new ID
    newList = append([]string{idStr}, newList...)

    // Trim to max length
    if len(newList) > maxItems {
        newList = newList[:maxItems]
    }

    // Set updated cookie
    http.SetCookie(w, &http.Cookie{
        Name:  cookieName,
        Value: strings.Join(newList, ","),
        Path:  "/",
    })
}

func getRecentlyViewedCars(r *http.Request, allData *Data) []CarModel {
    cookie, err := r.Cookie("recently_viewed")
    if err != nil {
        return nil
    }

    ids := strings.Split(cookie.Value, ",")
    var result []CarModel

    for _, idStr := range ids {
        id, err := strconv.Atoi(idStr)
        if err != nil {
            continue
        }

        for _, car := range allData.CarModels {
            if car.ID == id {
                result = append(result, car)
                break
            }
        }
    }

    return result
}