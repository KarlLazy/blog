package controller

import (
	"net/http"
)

// Router router
func Router() (mux *http.ServeMux) {
	mux = http.NewServeMux()

	mux.HandleFunc("/user/login", Login)
	mux.HandleFunc("/user/logout", Logout)

	return mux
}
