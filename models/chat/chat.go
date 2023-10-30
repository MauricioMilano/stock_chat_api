package chat_model

import (
	"time"

	"github.com/MauricioMilano/stock_app/config"
	chatroom_model "github.com/MauricioMilano/stock_app/models/chatroom"
	user_model "github.com/MauricioMilano/stock_app/models/user"

	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	Timestamp  time.Time               `json:"Timestamp"`
	Message    string                  `json:"Message"`
	UserId     uint                    `json:"UserId" gorm:"index"`
	User       user_model.User         `json:"User" gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
	ChatRoomId uint                    `json:"ChatRoomId" gorm:"index"`
	ChatRoom   chatroom_model.ChatRoom `json:"ChatRoom" gorm:"foreignKey:ChatRoomId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
}

func (cr *Chat) List(roomId uint, cht *[]Chat) *gorm.DB {
	db := config.GetDB()
	db = db.Where(Chat{ChatRoomId: roomId}).Preload("ChatRoom").Preload("User").Find(&cht).Limit(50).Order("created_at DESC")
	return db
}

func (c *Chat) Add() *gorm.DB {
	db := config.GetDB()
	c.Timestamp = time.Now()

	db = db.Create(&c)
	return db
}
