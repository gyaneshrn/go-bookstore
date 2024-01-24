package routes

import (
	"github.com/gorilla/mux"
	"github.com/gyaneshrn/go-bookstore/pkg/controllers"
)

var RegisterBookStore = func(router *mux.Router) {
	router.handleFunc("/book", controllers.CreateBook).Methods("POST")
	router.handleFunc("/book", controllers.GetBook).Methods("GET")
	router.handleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.handleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.handleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
