package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	error_model "github.com/MauricioMilano/stock_app/models/error"
	"github.com/MauricioMilano/stock_app/services/rabbitmq"
	"github.com/MauricioMilano/stock_app/services/websocket"
	error_utils "github.com/MauricioMilano/stock_app/utils/error"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

var RegisterWebsocketRoute = func(router *mux.Router) {
	roomMap := websocket.NewRoomMap()
	sb := router.PathPrefix("/v1").Subrouter()
	sb.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		jwtToken := r.URL.Query().Get("jwt")
		jwtSecret := os.Getenv("JWT_SECRET")
		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})

		if err != nil {
			handleWebsocketAuthenticationErr(w, err)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			handleWebsocketAuthenticationErr(w, err)
			return
		}

		serveWS(roomMap, w, r, claims)

	})
}

func serveWS(roomMap *websocket.RoomMap, w http.ResponseWriter, r *http.Request, claims jwt.MapClaims) {
	conn, err := websocket.Upgrade(w, r)
	error_utils.ErrorCheck(err)
	br := rabbitmq.GetRabbitMQBroker()
	room_id := r.URL.Query().Get("room_id")
	pool := roomMap.GetPool(room_id)
	client := &websocket.Client{
		Connection: conn,
		Pool:       pool,
		Name:       claims["UserName"].(string),
		UserID:     uint(claims["UserID"].(float64)),
	}

	fmt.Println("Websocket ready to accept connections")
	pool.Register <- client
	requestBody := make(chan []byte) // websocket.Message byte array channel
	go client.Read(requestBody)
	go br.ReadMessages(roomMap)
	go br.PublishMessage(requestBody)
}

func handleWebsocketAuthenticationErr(w http.ResponseWriter, err error) {
	log.Println("websocket error: ", err)
	w.WriteHeader(http.StatusUnauthorized)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	res := error_model.ErrorResponse{Message: err.Error(), Status: false, Code: http.StatusUnauthorized}
	data, err := json.Marshal(res)
	error_utils.ErrorCheck(err)
	w.Write(data)
}
