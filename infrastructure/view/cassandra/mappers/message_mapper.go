package mappers

import (
	"github.com/vulpes-ferrilata/chat-service/infrastructure/view/cassandra/entities"
	"github.com/vulpes-ferrilata/chat-service/view/models"
)

func ToMessageView(messageEntity *entities.Message) *models.Message {
	if messageEntity == nil {
		return nil
	}

	message := &models.Message{
		ID:     messageEntity.ID,
		RoomID: messageEntity.RoomID.ObjectID,
		UserID: messageEntity.UserID.ObjectID,
		Detail: messageEntity.Detail,
	}

	return message
}
