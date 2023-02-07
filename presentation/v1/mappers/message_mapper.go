package mappers

import (
	pb_models "github.com/vulpes-ferrilata/chat-service-proto/pb/models"
	"github.com/vulpes-ferrilata/chat-service/view/models"
)

var MessageMapper messageMapper = messageMapper{}

type messageMapper struct{}

func (m messageMapper) ToResponse(messageView *models.Message) (*pb_models.Message, error) {
	if messageView == nil {
		return nil, nil
	}

	return &pb_models.Message{
		ID:     messageView.ID.String(),
		RoomID: messageView.RoomID.Hex(),
		UserID: messageView.UserID.Hex(),
		Detail: messageView.Detail,
	}, nil
}
