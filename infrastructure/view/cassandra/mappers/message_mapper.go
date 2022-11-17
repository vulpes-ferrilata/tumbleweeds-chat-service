package mappers

import (
	"github.com/vulpes-ferrilata/chat-service/infrastructure/view/cassandra/entities"
	"github.com/vulpes-ferrilata/chat-service/view/models"
)

var MessageMapper messageMapper = messageMapper{}

type messageMapper struct{}

func (m messageMapper) ToView(messageEntity *entities.Message) (*models.Message, error) {
	if messageEntity == nil {
		return nil, nil
	}

	return &models.Message{
		ID:     messageEntity.ID,
		RoomID: messageEntity.RoomID.ObjectID,
		UserID: messageEntity.UserID.ObjectID,
		Detail: messageEntity.Detail,
	}, nil
}
