package chatroom_model

import (
	"github.com/MauricioMilano/stock_app/config"
	chatmessage_model "github.com/MauricioMilano/stock_app/models/chat_message"

	"gorm.io/gorm"
)

type ChatRoomMessagesRequest struct {
	RoomId uint `json:"roomId"`
}
type ChatRoomCreateRequest struct {
	Name string `json:"name"`
}

type ChatRoomsResponse struct {
	ChatRooms []ChatRoom `json:"chatRooms"`
}
type ChatRoomResponse struct {
	ChatRoom ChatRoom `json:"chatRoom"`
}

type ChatRoomMessagesResponse struct {
	Chats []chatmessage_model.ChatMessage `json:"chats"`
}
type ChatRoom struct {
	gorm.Model
	Name string `json:"name"`
}

func (cr *ChatRoom) Add() *gorm.DB {
	db := config.GetDB()
	db = db.Where(cr).FirstOrCreate(&cr)
	return db
}

func (cr *ChatRoom) List(cht *[]ChatRoom) *gorm.DB {
	db := config.GetDB()
	db = db.Order("id DESC LIMIT 50").Find(&cht)
	return db

}
