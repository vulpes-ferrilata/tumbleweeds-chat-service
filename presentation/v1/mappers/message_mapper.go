package mappers

import (
	"github.com/vulpes-ferrilata/chat-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/chat-service/view/models"
)

var MessageMapper messageMapper = messageMapper{}

type messageMapper struct{}

func (m messageMapper) ToResponse(messageView *models.Message) (*responses.Message, error) {
	if messageView == nil {
		return nil, nil
	}

	return &responses.Message{
		ID:     messageView.ID.String(),
		RoomID: messageView.RoomID.Hex(),
		UserID: messageView.UserID.Hex(),
		Detail: messageView.Detail,
	}, nil
}
