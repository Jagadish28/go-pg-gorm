package routes

import (
	"db-test/pkg/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

var RegisterBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/book/", controllers.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods(http.MethodGet)
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods(http.MethodDelete)

}
