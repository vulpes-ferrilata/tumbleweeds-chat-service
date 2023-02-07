package queries

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/chat-service/view/models"
	"github.com/vulpes-ferrilata/chat-service/view/projectors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FindMessagesByRoomIDQuery struct {
	RoomID string `validate:"omitempty,objectid"`
}

func NewFindMessagesByRoomIDHandler(messageProjector projectors.MessageProjector) *FindMessagesByRoomIDHandler {
	return &FindMessagesByRoomIDHandler{
		messageProjector: messageProjector,
	}
}

type FindMessagesByRoomIDHandler struct {
	messageProjector projectors.MessageProjector
}

func (f FindMessagesByRoomIDHandler) Handle(ctx context.Context, findMessagesByRoomID *FindMessagesByRoomIDQuery) ([]*models.Message, error) {
	roomID, _ := primitive.ObjectIDFromHex(findMessagesByRoomID.RoomID)

	messages, err := f.messageProjector.FindByRoomID(ctx, roomID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return messages, nil
}
