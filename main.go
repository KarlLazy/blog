package main

import (
	"fmt"
	"net/http"
	// "blog/controller"
)

func main() {
	start()

	http.ListenAndServe(":8080", nil)
}

func start() {
	http.HandleFunc("/index", index)

	// http.Handle("/user/", controller.Router())
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
