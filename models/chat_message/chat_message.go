package chatmessage_model

import "time"

type ChatMessage struct {
	ChatMessage   string    `json:"chatMessage"`
	ChatUserID    uint      `json:"chatUserId"`
	ChatUser      string    `json:"chatUser"`
	ChatRoomId    uint      `json:"chatRoomId"`
	ChatRoomName  string    `json:"chatRoomName"`
	ChatMessageTs time.Time `json:"ChatMessageTs"`
}
