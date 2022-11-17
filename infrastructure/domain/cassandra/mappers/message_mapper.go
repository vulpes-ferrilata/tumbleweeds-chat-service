package mappers

import (
	"github.com/vulpes-ferrilata/chat-service/domain/models"
	"github.com/vulpes-ferrilata/chat-service/infrastructure/domain/cassandra/entities"
)

var MessageMapper messageMapper = messageMapper{}

type messageMapper struct{}

func (m messageMapper) ToEntity(message *models.Message) (*entities.Message, error) {
	if message == nil {
		return nil, nil
	}

	return &entities.Message{
		ID:     message.GetID(),
		RoomID: message.GetRoomID(),
		UserID: message.GetUserID(),
		Detail: message.GetDetail(),
	}, nil
}
