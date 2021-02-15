package routes

import (
	"GoBackend/controller"
	"GoBackend/models"
	"net/http"
	"strings"
)

// RegisterUser - Registers the routes for Users and how to handle them
func RegisterUser() {
	http.HandleFunc("/user", index)
	http.HandleFunc("/user/register", register)
	http.HandleFunc("/user/login", login)
	http.HandleFunc("/user/", show)
}

func index(w http.ResponseWriter, r *http.Request) {
	controller.UserIndex(w, r)
}

func show(w http.ResponseWriter, r *http.Request) {
	param := strings.Split(r.URL.Path, "/")[2]
	if param == "" {
		index(w, r)
		return
	}

	user := models.FindUser(param)
	if user == nil {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case "GET":
		controller.UserShow(user, w, r)
	case "PUT":
		controller.UserUpdate(user, w, r)
	case "DELETE":
		controller.UserDestroy(user, w, r)
	}

}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		controller.ShowRegister(w, r)
		return
	}
	if r.Method == "POST" {
		controller.StoreRegister(w, r)
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		controller.ShowLogin(w, r)
		return
	}
	if r.Method == "POST" {
		controller.StoreLogin(w, r)
		return
	}
}
