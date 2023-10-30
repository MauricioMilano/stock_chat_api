package controllers

import (
	"encoding/json"
	"net/http"

	chatroom_model "github.com/MauricioMilano/stock_app/models/chatroom"
	"github.com/MauricioMilano/stock_app/services"
	"github.com/MauricioMilano/stock_app/utils"
	error_utils "github.com/MauricioMilano/stock_app/utils/error"
)

type ChatController struct {
	chatService services.Chat
}

func (c *ChatController) RegisterService(s services.Chat) {
	c.chatService = s
}

func (c *ChatController) ChatRooms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodPost || r.Method == http.MethodGet {
		res, err := c.chatService.ChatRooms()
		if err != nil {
			utils.ErrResponse(err, w)
			return
		}

		data, err := json.Marshal(res)
		error_utils.ErrorCheck(err)

		utils.Ok(data, w)
	}
}

func (c *ChatController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodPost {
		cP := chatroom_model.ChatRoomCreateRequest{}
		err := utils.ParseBody(r, &cP)
		if err != nil {
			utils.ErrResponse(error_utils.ErrInRequestMarshaling, w)
			return
		}

		res, err := c.chatService.CreateChatRoom(cP.Name)
		if err != nil {
			utils.ErrResponse(err, w)
			return
		}

		data, err := json.Marshal(res)
		error_utils.ErrorCheck(err)

		utils.Ok(data, w)
	}
}

func (c *ChatController) ChatRoomMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodPost {
		cmP := chatroom_model.ChatRoomMessagesRequest{}
		err := utils.ParseBody(r, &cmP)
		if err != nil {
			utils.ErrResponse(error_utils.ErrInRequestMarshaling, w)
			return
		}

		res, err := c.chatService.ChatRoomMessages(cmP.RoomId)
		if err != nil {
			utils.ErrResponse(err, w)
			return
		}

		data, err := json.Marshal(res)
		error_utils.ErrorCheck(err)

		utils.Ok(data, w)
	}
}
