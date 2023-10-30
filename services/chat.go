package services

import (
	"log"

	chat_model "github.com/MauricioMilano/stock_app/models/chat"
	chatmessage_model "github.com/MauricioMilano/stock_app/models/chat_message"
	chatroom_model "github.com/MauricioMilano/stock_app/models/chatroom"
	error_utils "github.com/MauricioMilano/stock_app/utils/error"
)

type Chat interface {
	CreateChatRoom(name string) (chatroom_model.ChatRoomResponse, error)
	ChatRooms() (chatroom_model.ChatRoomsResponse, error)
	ChatRoomMessages(roomId uint) (chatroom_model.ChatRoomMessagesResponse, error)
}

type ChatSaver interface {
	SaveChatMessage(msg string, roomId, userId uint) bool
}

type chat struct{}

func NewChatService() *chat {
	return &chat{}
}

func (c *chat) ChatRooms() (chatroom_model.ChatRoomsResponse, error) {
	var chtList []chatroom_model.ChatRoom
	var chtRoomModel chatroom_model.ChatRoom
	err := chtRoomModel.List(&chtList)
	error_utils.DBErrorCheck(err)
	return chatroom_model.ChatRoomsResponse{ChatRooms: chtList}, nil
}

func (c *chat) CreateChatRoom(name string) (chatroom_model.ChatRoomResponse, error) {
	cht := chatroom_model.ChatRoom{
		Name: name,
	}
	err := cht.Add()
	error_utils.DBErrorCheck(err)
	return chatroom_model.ChatRoomResponse{ChatRoom: cht}, nil
}

func (c *chat) ChatRoomMessages(roomId uint) (chatroom_model.ChatRoomMessagesResponse, error) {
	var chtList []chat_model.Chat
	chtMsgList := []chatmessage_model.ChatMessage{}
	var chtModel chat_model.Chat
	err := chtModel.List(roomId, &chtList)
	error_utils.DBErrorCheck(err)
	for _, v := range chtList {
		chtMsgList = append(chtMsgList, chatmessage_model.ChatMessage{
			ChatMessage:   v.Message,
			ChatUser:      v.User.UserName,
			ChatUserID:    v.User.ID,
			ChatMessageTs: v.Timestamp,
			ChatRoomId:    v.ChatRoomId,
			ChatRoomName:  v.ChatRoom.Name,
		})
	}
	return chatroom_model.ChatRoomMessagesResponse{Chats: chtMsgList}, nil
}

func (c *chat) SaveChatMessage(msg string, roomId uint, userId uint) bool {
	cht := chat_model.Chat{
		Message:    msg,
		UserId:     userId,
		ChatRoomId: roomId,
	}
	err := cht.Add()
	if err.Error != nil {
		log.Println("error: ", err.Error)
		return false
	}

	return true
}
