package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", helloWorld)
	fmt.Println("working on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
