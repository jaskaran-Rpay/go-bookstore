package v1

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/jaskaran-Rpay/go-bookstore/pkg/controllers"
)

var RegisterV1Routes = func(router *mux.Router) {
	fmt.Println("Registering V1 Route")
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
}
