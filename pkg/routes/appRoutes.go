package routes

import (
	"github.com/gorilla/mux"
	"github.com/oluwafemi08/go-app/pkg/controllers"
)

var RegisterAppRoutes = func(router *mux.Router) {
	router.HandleFunc("/user/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/user/{userID}", controllers.GetUserByID).Methods("GET")
	router.HandleFunc("/user/{userID}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{userID}", controllers.DeleteUser).Methods("DELETE")

}
