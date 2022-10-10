package servers

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/chat-service-proto/pb"
	"github.com/vulpes-ferrilata/chat-service-proto/pb/requests"
	"github.com/vulpes-ferrilata/chat-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/chat-service/application/commands"
	"github.com/vulpes-ferrilata/chat-service/application/queries"
	"github.com/vulpes-ferrilata/chat-service/infrastructure/cqrs/command"
	"github.com/vulpes-ferrilata/chat-service/infrastructure/cqrs/query"
	"github.com/vulpes-ferrilata/chat-service/infrastructure/utils/slices"
	"github.com/vulpes-ferrilata/chat-service/presentation/v1/mappers"
	"github.com/vulpes-ferrilata/chat-service/view/models"
	"google.golang.org/protobuf/types/known/emptypb"
)

func NewChatServer(findMessagesByRoomIDQueryHandler query.QueryHandler[*queries.FindMessagesByRoomID, []*models.Message],
	createMessageCommandHandler command.CommandHandler[*commands.CreateMessage]) pb.ChatServer {
	return &chatServer{
		findMessagesByRoomIDQueryHandler: findMessagesByRoomIDQueryHandler,
		createMessageCommandHandler:      createMessageCommandHandler,
	}
}

type chatServer struct {
	pb.UnimplementedChatServer
	findMessagesByRoomIDQueryHandler query.QueryHandler[*queries.FindMessagesByRoomID, []*models.Message]
	createMessageCommandHandler      command.CommandHandler[*commands.CreateMessage]
}

func (c chatServer) FindMessagesByRoomID(ctx context.Context, findMessagesByRoomIDRequest *requests.FindMessagesByRoomID) (*responses.MessageList, error) {
	findMessagesByRoomIDQuery := &queries.FindMessagesByRoomID{
		RoomID: findMessagesByRoomIDRequest.GetRoomID(),
	}

	messages, err := c.findMessagesByRoomIDQueryHandler.Handle(ctx, findMessagesByRoomIDQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	messageResponses, err := slices.Map(func(message *models.Message) (*responses.Message, error) {
		return mappers.ToMessageResponse(message), nil
	}, messages)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	messageListResponse := &responses.MessageList{
		Messages: messageResponses,
	}

	return messageListResponse, nil
}

func (c chatServer) CreateMessage(ctx context.Context, createMessageRequest *requests.CreateMessage) (*emptypb.Empty, error) {
	createMessageCommand := &commands.CreateMessage{
		MessageID: createMessageRequest.GetMessageID(),
		RoomID:    createMessageRequest.GetRoomID(),
		UserID:    createMessageRequest.GetUserID(),
		Detail:    createMessageRequest.GetDetail(),
	}

	if err := c.createMessageCommandHandler.Handle(ctx, createMessageCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}
