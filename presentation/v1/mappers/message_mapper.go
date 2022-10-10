package mappers

import (
	"github.com/vulpes-ferrilata/chat-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/chat-service/view/models"
)

func ToMessageResponse(messageView *models.Message) *responses.Message {
	if messageView == nil {
		return nil
	}

	return &responses.Message{
		ID:     messageView.ID.String(),
		RoomID: messageView.RoomID.Hex(),
		UserID: messageView.UserID.Hex(),
		Detail: messageView.Detail,
	}
}
