package routes

import (
	"github.com/MauricioMilano/stock_app/config"
	"github.com/MauricioMilano/stock_app/controllers"
	"github.com/MauricioMilano/stock_app/middlewares"
	"github.com/MauricioMilano/stock_app/services"

	"github.com/gorilla/mux"
)

var RegisterChatRoutes = func(router *mux.Router, opts config.ConfigOpts) {

	sb := router.PathPrefix("/v1/api/chat").Subrouter()
	sb.Use(middlewares.HeaderMiddleware)
	sb.Use(middlewares.Authenticated)

	var chat controllers.ChatController
	chat.RegisterService(services.NewChatService())

	sb.HandleFunc("/create", chat.Create).Methods("POST", "OPTIONS")
	sb.HandleFunc("/rooms", chat.ChatRooms).Methods("POST", "GET", "OPTIONS")
	sb.HandleFunc("/room-messages", chat.ChatRoomMessages).Methods("POST", "OPTIONS")
}
