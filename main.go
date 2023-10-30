package main

import (
	"errors"
	"fmt"
	"log"
	stdlog "log"
	"net/http"
	"os"

	"github.com/MauricioMilano/stock_app/config"
	"github.com/MauricioMilano/stock_app/middlewares"
	chat_model "github.com/MauricioMilano/stock_app/models/chat"
	chatroom_model "github.com/MauricioMilano/stock_app/models/chatroom"
	user_model "github.com/MauricioMilano/stock_app/models/user"
	"github.com/MauricioMilano/stock_app/routes"
	"github.com/MauricioMilano/stock_app/services/rabbitmq"
	error_utils "github.com/MauricioMilano/stock_app/utils/error"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

}

func run() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Connect Rabbit MQ
	conn, ch := rabbitmq.InitilizeBroker()
	defer conn.Close()
	defer ch.Close()

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		stdlog.Println("JWT Secret not set")
		return errors.New("JWT Secret not set")
	}
	opts := config.ConfigOpts{}
	opts.ConnectDB()
	db := config.GetDB()
	err = db.AutoMigrate(&user_model.User{}, &chatroom_model.ChatRoom{}, &chat_model.Chat{})
	error_utils.ErrorCheck(err)

	r := mux.NewRouter()

	routes.RegisterAuthRoutes(r, opts)
	routes.RegisterChatRoutes(r, opts)
	routes.RegisterWebsocketRoute(r)
	handler := middlewares.EnableCors(r)

	// Start api server
	port := os.Getenv("APP_PORT")
	fmt.Println("App started")

	err = http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
	return err
}
