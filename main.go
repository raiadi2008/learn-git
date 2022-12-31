package main

import (
	"log"
	"net/http"
)

const port = ":3000"

func Home(w http.ResponseWriter, r *http.Request) {
	// this code is changes
}

func SumNum(a, b int) int {
	return a + b
}

func multnum(a, b int) int {
	return a * b
}

func About(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Printf("App started on port http://127.0.0.1%s", port)
	_ = http.ListenAndServe(port, nil)
}
