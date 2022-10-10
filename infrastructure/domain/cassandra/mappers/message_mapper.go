package mappers

import (
	"github.com/vulpes-ferrilata/chat-service/domain/models"
	"github.com/vulpes-ferrilata/chat-service/infrastructure/domain/cassandra/entities"
)

func ToMessageEntity(message *models.Message) *entities.Message {
	if message == nil {
		return nil
	}

	return &entities.Message{
		ID:     message.GetID(),
		RoomID: message.GetRoomID(),
		UserID: message.GetUserID(),
		Detail: message.GetDetail(),
	}
}
