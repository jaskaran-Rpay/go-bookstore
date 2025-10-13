package routes

import (
	"github.com/gorilla/mux"
	"github.com/jaskaran-Rpay/go-bookstore/pkg/controllers"
	v1 "github.com/jaskaran-Rpay/go-bookstore/pkg/routes/v1"
)

// added comment

var RegisterBookStoreRoutes = func(router *mux.Router) {

	v1.RegisterV1Routes(router.PathPrefix("/v1").Subrouter())

	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}/{moreid}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
