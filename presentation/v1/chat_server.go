package v1

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/chat-service-proto/pb"
	pb_models "github.com/vulpes-ferrilata/chat-service-proto/pb/models"
	"github.com/vulpes-ferrilata/chat-service/application/commands"
	"github.com/vulpes-ferrilata/chat-service/application/queries"
	"github.com/vulpes-ferrilata/chat-service/presentation/v1/mappers"
	"github.com/vulpes-ferrilata/chat-service/view/models"
	"github.com/vulpes-ferrilata/cqrs"
	"github.com/vulpes-ferrilata/slices"
	"google.golang.org/protobuf/types/known/emptypb"
)

func NewChatServer(queryBus *cqrs.QueryBus,
	commandBus *cqrs.CommandBus) pb.ChatServer {
	return &chatServer{
		queryBus:   queryBus,
		commandBus: commandBus,
	}
}

type chatServer struct {
	pb.UnimplementedChatServer
	queryBus   *cqrs.QueryBus
	commandBus *cqrs.CommandBus
}

func (c chatServer) FindMessagesByRoomID(ctx context.Context, findMessagesByRoomIDRequest *pb_models.FindMessagesByRoomIDRequest) (*pb_models.MessageList, error) {
	findMessagesByRoomIDQuery := &queries.FindMessagesByRoomIDQuery{
		RoomID: findMessagesByRoomIDRequest.GetRoomID(),
	}

	messages, err := cqrs.ParseQueryHandlerFunc[*queries.FindMessagesByRoomIDQuery, []*models.Message](c.queryBus.Execute)(ctx, findMessagesByRoomIDQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	messageResponses, err := slices.Map(func(message *models.Message) (*pb_models.Message, error) {
		return mappers.MessageMapper.ToResponse(message)
	}, messages...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	messageListResponse := &pb_models.MessageList{
		Messages: messageResponses,
	}

	return messageListResponse, nil
}

func (c chatServer) CreateMessage(ctx context.Context, createMessageRequest *pb_models.CreateMessageRequest) (*emptypb.Empty, error) {
	createMessageCommand := &commands.CreateMessageCommand{
		MessageID: createMessageRequest.GetMessageID(),
		RoomID:    createMessageRequest.GetRoomID(),
		UserID:    createMessageRequest.GetUserID(),
		Detail:    createMessageRequest.GetDetail(),
	}

	if err := c.commandBus.Execute(ctx, createMessageCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}
