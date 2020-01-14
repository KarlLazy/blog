package controller

import (
	"net/http"
)

// Login user login
func Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login success"))
}

// Logout user logout
func Logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("logout success"))
}
