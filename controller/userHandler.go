package controller

import (
	"blog/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Login user login
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		userName := r.FormValue("userName")
		password := r.FormValue("pwd")

		user := &model.User{
			UserName: userName,
			Password: password,
		}

		fmt.Fprintf(w, "userName: %s, password: %s\n", userName, password)

		User.LoginGet(user)
		return
	} else if r.Method == "POST" {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Read all body error: %s", err)
			return
		}

		user := new(model.User)
		err = json.Unmarshal(data, user)
		if err != nil {
			fmt.Printf("Unmarshal body error: %s", err)
			return
		}

		fmt.Fprintf(w, "userName: %s, password: %s\n", user.UserName, user.Password)

		User.LoginPost(user)
	}

}

// Logout user login
func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("user logout")

	w.Write([]byte("logout success"))
}
