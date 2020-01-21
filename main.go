package main

import (
	"fmt"
	"net/http"

	"blog/config"
	"blog/controller"
)

func main() {
	start()

	http.ListenAndServe(config.Config.HTTP.Addr, nil)
}

func start() {
	http.HandleFunc("/index", index)

	http.Handle("/user/", controller.Router())
	http.Handle("/order/", controller.Router())
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in")
	fmt.Fprintf(w, "Hello World")
}
