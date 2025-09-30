package lib

import (
	"fmt"
	"strconv"
)

func renderCarList(allData, data *Data, recentCars []CarModel, selectedFilters map[string]string) string {
	allYears := make(map[int]bool)
	allHorsepower := make(map[int]bool)
	allEngines := make(map[string]bool)

	for _, car := range allData.CarModels {
		allYears[car.Year] = true
		allHorsepower[car.Specifications.Horsepower] = true
		allEngines[car.Specifications.Engine] = true
	}

	type HorsepowerRange struct {
		Label    string
		Min, Max int
		Value    string // used in `<option value="">`
	}

	var horsepowerRanges = []HorsepowerRange{
		{"Under 100 HP", 0, 99, "0-99"},
		{"100-199 HP", 100, 199, "100-199"},
		{"200-299 HP", 200, 299, "200-299"},
		{"300-399 HP", 300, 399, "300-399"},
		{"400 HP and above", 400, 10000, "400-10000"},
	}

	html := `<html>
	<head>
		<title>Car Dealership</title>
		<link rel="stylesheet" href="/static/Main.css">
	</head>
	<body>
		<nav class="navbar">
        	<a href="/" class="nav-button">Home</a>
        	<a href="/compare" class="nav-button">Compare</a>
   		</nav>
		<div class="container">
			
			<aside class="filters">`
	html += fmt.Sprintf(`
			<h2>Search Cars</h2>
			<form action="/" method="GET">
				  <input type="text" name="search" placeholder="Search cars..." value="%s" />
				  <button type="submit">Search</button>
			</form>`, selectedFilters["search"])
	html += `<h2>Filter Cars</h2>
				<form class="filter-form" method="GET" action="/">
					<label for="manufacturer">Manufacturer:</label>
					<select name="manufacturer">
						<option value="">All</option>`
	for _, m := range allData.Manufacturers {
		selected := ""
		if selectedFilters["manufacturer"] == fmt.Sprintf("%d", m.ID) {
			selected = "selected"
		}
		html += fmt.Sprintf(`<option value="%d" %s>%s</option>`, m.ID, selected, m.Name)
	}
	html += `</select>

					<label for="category">Category:</label>
					<select name="category">
						<option value="">All</option>`
	for _, c := range allData.Categories {
		selected := ""
		if selectedFilters["category"] == fmt.Sprintf("%d", c.ID) {
			selected = "selected"
		}
		html += fmt.Sprintf(`<option value="%d" %s>%s</option>`, c.ID, selected, c.Name)
	}
	html += `</select>

					<label for="year">Year:</label>
					<select name="year">
						<option value="">All</option>`
	for year := range allYears {
		selected := ""
		if selectedFilters["year"] == fmt.Sprintf("%d", year) {
			selected = "selected"
		}
		html += fmt.Sprintf(`<option value="%d" %s>%d</option>`, year, selected, year)
	}
	html += `</select>

					<label for="horsepower">Horsepower:</label>
					<select name="horsepower">
						<option value="">All</option>`
	for _, r := range horsepowerRanges {
		selected := ""
		if selectedFilters["horsepower"] == r.Value {
			selected = "selected"
		}
		html += fmt.Sprintf(`<option value="%s" %s>%s</option>`, r.Value, selected, r.Label)
	}
	html += `</select>

					<label for="engine">Engine:</label>
					<select name="engine">
						<option value="">All</option>`
	for engine := range allEngines {
		selected := ""
		if selectedFilters["engine"] == engine {
			selected = "selected"
		}
		html += fmt.Sprintf(`<option value="%s" %s>%s</option>`, engine, selected, engine)
	}
	html += `</select>

					<button type="submit">Apply</button>
					<button type="button" onclick="window.location='/'">Reset</button>
				</form>
			</aside>

			<div class="car-list">`

	for _, car := range data.CarModels {
		imageURL := fmt.Sprintf("http://localhost:3000/api/images/%s", car.Image)
		html += fmt.Sprintf(`
			<div class="car-item">
				<a href="/car?id=%d">
					<img src="%s" class="car-image">
					<br>
					<span class="car-name">%s</span>
				</a>
			</div>`, car.ID, imageURL, car.Name)
	}

	html += `</div>`

	if recentCars != nil {
		html += `<aside class="recently-viewed">
				<h2>Recently Viewed</h2>`

		for _, car := range recentCars {
			imageURL := fmt.Sprintf("http://localhost:3000/api/images/%s", car.Image)
			html += fmt.Sprintf(`
			<div class="recent-car-item">
				<a href="/car?id=%d">
					<img src="%s" class="recent-car-image">
					<span class="recent-car-name">%s</span>
				</a>
			</div>`, car.ID, imageURL, car.Name)
		}
	}

	html += `</div></body></html>`
	return html
}

func renderCarDetails(car CarModel, manufacturers []Manufacturer, categories []Category) string {
	var manufacturerName, Country, foundingYear, categoryName string

	// Find manufacturer name and founding year
	for _, m := range manufacturers {
		if m.ID == car.ManufacturerID {
			manufacturerName = m.Name
			foundingYear = strconv.Itoa(m.FoundingYear)
			Country = m.Country
			break
		}
	}

	// Find category name
	for _, c := range categories {
		if c.ID == car.CategoryID {
			categoryName = c.Name
			break
		}
	}

	imageURL := fmt.Sprintf("http://localhost:3000/api/images/%s", car.Image)

	html := `<html>
	<head>
		<title>` + car.Name + ` - Details</title>
		<link rel="stylesheet" href="/static/Main.css">
	</head>
	<body>
		<nav class="navbar">
			<a href="/" class="back-link">← Back to Cars</a>
        	<a href="/" class="nav-button">Home</a>
        	<a href="/compare" class="nav-button">Compare</a>
   		</nav>
		
		<h1>` + car.Name + `</h1>
		<div class="car-details">
			<img src="` + imageURL + `" class="car-image-large">
			<div class="comparison-table">
				<table>
					<tr><th>Manufacturer</th><th>` + manufacturerName + `</th></tr>
					<tr><td>Country</td><td>` + Country + `</td></tr>
					<tr><td>Founded</td><td>` + foundingYear + `</td></tr>
					<tr><td>Category</td><td>` + categoryName + `</td></tr>
					<tr><td>Year</td><td>` + fmt.Sprintf("%d", car.Year) + `</td></tr>
					<tr><td>Engine</td><td>` + car.Specifications.Engine + `</td></tr>
					<tr><td>Horsepower</td><td>` + fmt.Sprintf("%d HP", car.Specifications.Horsepower) + `</td></tr>
					<tr><td>Transmission</td><td>` + car.Specifications.Transmission + `</td></tr>
					<tr><td>Drivetrain</td><td>` + car.Specifications.Drivetrain + `</td></tr>
				</table>
			</div>
		</div>
	</body>
	</html>`

	return html
}

func renderComparisonPage(cars []CarModel, manufacturers []Manufacturer, categories []Category, car1, car2 *CarModel) string {
	var manufacturerName1, Country1, foundingYear1, categoryName1 string
	var manufacturerName2, Country2, foundingYear2, categoryName2 string

	html := `<html>
	<head>
		<title>Compare Cars</title>
		<link rel="stylesheet" href="/static/Main.css">
	</head>
	<body>
		<div class="navbar">
			<a href="/" class="back-link">← Back to Cars</a>
			<a href="/" class="nav-button nav-left">Home</a>
			<a href="/compare" class="nav-button">Compare</a>
		</div>

		<h1>Compare Two Cars</h1>
		<form method="GET" action="/compare">
			<select class="compare-select" name="car1" required>
				<option value="">Select Car 1</option>`
	for _, car := range cars {
		selected := ""
		if car1 != nil && car.ID == car1.ID {
			selected = "selected"
		}
		html += fmt.Sprintf(`<option value="%d" %s>%s</option>`, car.ID, selected, car.Name)
	}
	html += `</select>

			<select class="compare-select" name="car2" required>
				<option value="">Select Car 2</option>`
	for _, car := range cars {
		selected := ""
		if car2 != nil && car.ID == car2.ID {
			selected = "selected"
		}
		html += fmt.Sprintf(`<option value="%d" %s>%s</option>`, car.ID, selected, car.Name)
	}
	html += `</select>

			<button type="submit" class="nav-button">Submit</button>
		</form>`

	if car1 != nil && car2 != nil {
		// Find manufacturer name and founding year
		for _, m := range manufacturers {
			if m.ID == car1.ManufacturerID {
				manufacturerName1 = m.Name
				foundingYear1 = strconv.Itoa(m.FoundingYear)
				Country1 = m.Country
			}
			if m.ID == car2.ManufacturerID {
				manufacturerName2 = m.Name
				foundingYear2 = strconv.Itoa(m.FoundingYear)
				Country2 = m.Country
			}
		}

		// Find category name
		for _, c := range categories {
			if c.ID == car1.CategoryID {
				categoryName1 = c.Name
			}
			if c.ID == car2.CategoryID {
				categoryName2 = c.Name
			}
		}

		html += `
		<div class="comparison-table">
			<table>
				<tr><th>Specification</th><th>` + car1.Name + `</th><th>` + car2.Name + `</th></tr>
				<tr><td>Manufacturer</td><td>` + manufacturerName1 + `</td><td>` + manufacturerName2 + `</td></tr>
				<tr><td>Country</td><td>` + Country1 + `</td><td>` + Country2 + `</td></tr>
				<tr><td>Founded</td><td>` + foundingYear1 + `</td><td>` + foundingYear2 + `</td></tr>
				<tr><td>Category</td><td>` + categoryName1 + `</td><td>` + categoryName2 + `</td></tr>
				<tr><td>Year</td><td>` + fmt.Sprint(car1.Year) + `</td><td>` + fmt.Sprint(car2.Year) + `</td></tr>
				<tr><td>Engine</td><td>` + car1.Specifications.Engine + `</td><td>` + car2.Specifications.Engine + `</td></tr>
				<tr><td>Horsepower</td><td>` + fmt.Sprintf("%d HP", car1.Specifications.Horsepower) + `</td><td>` + fmt.Sprintf("%d HP", car2.Specifications.Horsepower) + `</td></tr>
				<tr><td>Transmission</td><td>` + car1.Specifications.Transmission + `</td><td>` + car2.Specifications.Transmission + `</td></tr>
				<tr><td>Drivetrain</td><td>` + car1.Specifications.Drivetrain + `</td><td>` + car2.Specifications.Drivetrain + `</td></tr>
			</table>
		</div>`
	}

	html += `</body></html>`
	return html
}

func render500Page() string {
	return `<html>
	<link rel="stylesheet" href="/static/Main.css">
	<head><title>Server Error</title></head>
	<body>
		<h1>500 - Internal Server Error</h1>
		<h2>Oops! Something went wrong on our side.</h2>
		<button type="button" class="nav-button" onclick="window.location='/'">Back to Home</button>
	</body>
</html>`
}
