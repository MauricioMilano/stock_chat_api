package routes

import (
	"github.com/MauricioMilano/stock_app/config"
	"github.com/MauricioMilano/stock_app/controllers"
	"github.com/MauricioMilano/stock_app/middlewares"
	"github.com/MauricioMilano/stock_app/services"

	"github.com/gorilla/mux"
)

var RegisterAuthRoutes = func(router *mux.Router, config config.ConfigOpts) {

	sb := router.PathPrefix("/v1/api/auth").Subrouter()
	sb.Use(middlewares.HeaderMiddleware)
	var auth controllers.AuthController
	auth.RegisterService(services.NewAuthService())
	sb.HandleFunc("/login", auth.Login).Methods("POST", "OPTIONS")
	sb.HandleFunc("/signup", auth.SignUp).Methods("POST", "OPTIONS")
}
