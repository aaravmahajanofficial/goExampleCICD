package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello From Jenkins CI/CD!!!!!!!!!")

}

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":9090", nil)

}
