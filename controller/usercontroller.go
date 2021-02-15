package controller

import (
	"GoBackend/models"
	"fmt"
	"net/http"
)

// UserIndex shows all users
func UserIndex(w http.ResponseWriter, r *http.Request) {
	users := models.AllUsers()
	for i := 0; i < len(users); i++ {
		fmt.Fprintln(w, users[i].Name)
	}
}

// UserShow shows one user
func UserShow(user *models.User, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, user.Name)
}

// UserUpdate updates the user
func UserUpdate(user *models.User, w http.ResponseWriter, r *http.Request) {

}

// UserDestroy destroys the user
func UserDestroy(user *models.User, w http.ResponseWriter, r *http.Request) {

}
