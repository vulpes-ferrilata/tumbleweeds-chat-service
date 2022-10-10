package handlers

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gocql/gocql"
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/chat-service/application/commands"
	"github.com/vulpes-ferrilata/chat-service/domain/models"
	"github.com/vulpes-ferrilata/chat-service/domain/repositories"
	"github.com/vulpes-ferrilata/chat-service/infrastructure/cassandra/udt"
	"github.com/vulpes-ferrilata/chat-service/infrastructure/cqrs/command"
	"github.com/vulpes-ferrilata/chat-service/infrastructure/cqrs/command/wrappers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewCreateMessageCommandHandler(validate *validator.Validate, messageRepository repositories.MessageRepository) command.CommandHandler[*commands.CreateMessage] {
	handler := &createMessageCommandHandler{
		messageRepository: messageRepository,
	}
	validationWrapper := wrappers.NewValidationWrapper[*commands.CreateMessage](validate, handler)

	return validationWrapper
}

type createMessageCommandHandler struct {
	messageRepository repositories.MessageRepository
}

func (c createMessageCommandHandler) Handle(ctx context.Context, createMessageCommand *commands.CreateMessage) error {
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
