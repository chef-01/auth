package router

import (
	"auth/modules/user/presentation/controllers"
	"github.com/gorilla/mux"
)

/*
SetupUserRoutes configures the HTTP routes for user-related operations.

Params:
- router: The main Gorilla Mux router.
- controller: UserController that handles incoming HTTP requests.

Usage:
Call this function from your DI container to register all user routes.
*/
func SetupUserRoutes(router *mux.Router, controller *controllers.UserController) {
	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("", controller.CreateUser).Methods("POST")
	userRouter.HandleFunc("", controller.GetAllUsers).Methods("GET")
	userRouter.HandleFunc("/{id}", controller.UpdateUser).Methods("PUT")
}
