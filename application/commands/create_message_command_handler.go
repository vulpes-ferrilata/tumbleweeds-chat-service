package commands

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/chat-service/domain/models"
	"github.com/vulpes-ferrilata/chat-service/domain/repositories"
	"github.com/vulpes-ferrilata/chat-service/infrastructure/cassandra/udt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateMessageCommand struct {
	MessageID string `validate:"required,uuid"`
	RoomID    string `validate:"required,objectid"`
	UserID    string `validate:"required,objectid"`
	Detail    string `validate:"required"`
}

func NewCreateMessageCommandHandler(messageRepository repositories.MessageRepository) *CreateMessageCommandHandler {
	return &CreateMessageCommandHandler{
		messageRepository: messageRepository,
	}
}

type CreateMessageCommandHandler struct {
	messageRepository repositories.MessageRepository
}

func (c CreateMessageCommandHandler) Handle(ctx context.Context, createMessageCommand *CreateMessageCommand) error {
	messageID, err := gocql.ParseUUID(createMessageCommand.MessageID)
	if err != nil {
		return errors.WithStack(err)
	}

	roomID, err := primitive.ObjectIDFromHex(createMessageCommand.RoomID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(createMessageCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	message := models.MessageBuilder{}.
		SetID(messageID).
		SetRoomID(udt.ObjectID{roomID}).
		SetUserID(udt.ObjectID{userID}).
		SetDetail(createMessageCommand.Detail).
		Create()

	if err := c.messageRepository.Insert(ctx, message); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
