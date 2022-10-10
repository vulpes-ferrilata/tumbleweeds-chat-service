package handlers

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/chat-service/application/queries"
	"github.com/vulpes-ferrilata/chat-service/infrastructure/cqrs/query"
	"github.com/vulpes-ferrilata/chat-service/infrastructure/cqrs/query/wrappers"
	"github.com/vulpes-ferrilata/chat-service/view/models"
	"github.com/vulpes-ferrilata/chat-service/view/projectors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewFindMessagesByRoomIDHandler(validate *validator.Validate, messageProjector projectors.MessageProjector) query.QueryHandler[*queries.FindMessagesByRoomID, []*models.Message] {
	handler := &findMessagesByRoomIDHandler{
		messageProjector: messageProjector,
	}
	validationWrapper := wrappers.NewValidationWrapper[*queries.FindMessagesByRoomID, []*models.Message](validate, handler)

	return validationWrapper
}

type findMessagesByRoomIDHandler struct {
	messageProjector projectors.MessageProjector
}

func (f findMessagesByRoomIDHandler) Handle(ctx context.Context, findMessagesByRoomID *queries.FindMessagesByRoomID) ([]*models.Message, error) {
	roomID, _ := primitive.ObjectIDFromHex(findMessagesByRoomID.RoomID)

	messages, err := f.messageProjector.FindByRoomID(ctx, roomID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return messages, nil
}
