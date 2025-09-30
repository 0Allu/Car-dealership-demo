# Car Dealership Web Server (Go)

This is a Go web server for a car dealership website. It displays car listings, allows filtering and searching, viewing detailed car information, and comparing cars side-by-side. It fetches data from a local API and renders HTML pages.

---

## 🚀 Features

- 🏎️ List all car models from the API  
- 🔍 Filter cars by:
  - Manufacturer  
  - Category  
  - Year  
  - Horsepower range  
  - Engine type  
  - Text search 
- 📄 View detailed information for each car  
- 🆚 Compare two cars side-by-side  
- 🕓 Track and display recently viewed cars using cookies  

---

## 📁 Project Structure

```
cars/
├── api/                    # Node.js + Express API with data.json
├── lib/
│   ├── fetching.go         # Data fetching logic from the API
│   ├── handles.go          # HTTP handlers (main page, car page, comparison)
│   ├── renders.go          # HTML rendering functions
│   └── structs.go          # Contains shared data structures used across the server, including:
│                           # - Manufacturer: represents a car brand (e.g., Toyota, BMW)
│                           # - Category: represents a car type (e.g., SUV, Sedan)
│                           # - CarModel: represents an individual car with detailed specs
│                           # - Data: wrapper struct holding all API data for easy access
├── static/                 # Static assets (CSS, images, etc.)
├── main.go                 # Entry point of the Go server
└── README.md
```

---

## 🌐 How It Works

1. `fetchData()` makes a request to the API at `http://localhost:3000/api` to get the list of endpoint URLs.  
2. It then concurrently fetches:
   - Manufacturers  
   - Categories  
   - Car models  
3. Results are rendered using a manually constructed HTML.  
4. Query parameters are read and used to filter data on the `/` route. (main page)
5. Individual car pages are shown using the `/car?id=X` route.  (car details)
6. Cars can be compared via the `/compare?car1=ID1&car2=ID2` route. (compare page)
7. Recently viewed cars are tracked using cookies and displayed on the main page.  

---

## 🔧 Setup & Running

### 1. Start the API

The server expects a local API running on `http://localhost:3000`.

- Install [NodeJS](http://nodejs.org)
- Install [NPM](https://www.npmjs.com/package/npm) package manager

```bash
cd api
make build
make run
```

### 2. Run the Go Server

Start a new terminal

Then run:

```bash
go build cars
./cars
```

Then open your browser and go to:  
👉 `http://localhost:8080`

---

## 🔎 Example URLs

- `/` — main page  
- `/?manufacturer=2&category=1&year=2023` — filter by manufacturer, category, and year  
- `/?search=civic` — search by text  
- `/car?id=4` — view car details  
- `/compare?car1=5&car2=9` — compare two cars  

---

## 📌 Dependencies

- Go standard library only  
- No external packages  

---