package controller

import (
	"fmt"
	"net/http"
)

// ShowRegister renders the register page
func ShowRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register Page Reached")
}

// StoreRegister stores and attempts to register
func StoreRegister(w http.ResponseWriter, r *http.Request) {

}
