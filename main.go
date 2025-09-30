package main

import (
	"fmt"
	"net/http"
	"cars/lib"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", lib.MainPageHandler)
	http.HandleFunc("/car", lib.CarHandler)
	http.HandleFunc("/compare", lib.ComparePageHandler)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}