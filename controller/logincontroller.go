package controller

import (
	"fmt"
	"net/http"
)

// ShowLogin renders the login page
func ShowLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login Page Reached")
}

// StoreLogin stores and attempts to login
func StoreLogin(w http.ResponseWriter, r *http.Request) {

}
