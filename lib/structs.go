package lib

type Manufacturer struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Country      string `json:"country"`
	FoundingYear int    `json:"foundingYear"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Specifications struct {
	Engine       string `json:"engine"`
	Horsepower   int    `json:"horsepower"`
	Transmission string `json:"transmission"`
	Drivetrain   string `json:"drivetrain"`
}

type CarModel struct {
	ID            int            `json:"id"`
	Name          string         `json:"name"`
	ManufacturerID int           `json:"manufacturerId"`
	CategoryID    int            `json:"categoryId"`
	Year          int            `json:"year"`
	Specifications Specifications `json:"specifications"`
	Image         string         `json:"image"`
}

type Data struct {
	Manufacturers []Manufacturer `json:"manufacturers"`
	Categories    []Category     `json:"categories"`
	CarModels     []CarModel     `json:"carModels"`
}
