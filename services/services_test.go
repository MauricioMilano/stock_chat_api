package services

import (
	"testing"

	"github.com/MauricioMilano/stock_app/utils"

	"github.com/stretchr/testify/assert"
)

func TestChatService(t *testing.T) {
	utils.TestHelper()

	service := NewChatService()
	t.Run("ChatRooms Test", func(t *testing.T) {
		_, err := service.ChatRooms()
		assert.Equal(t, err, nil)
	})
}
