package infrastructure

import (
	command_handlers "github.com/vulpes-ferrilata/chat-service/application/commands/handlers"
	query_handlers "github.com/vulpes-ferrilata/chat-service/application/queries/handlers"
	"github.com/vulpes-ferrilata/chat-service/infrastructure/domain/cassandra/repositories"
	"github.com/vulpes-ferrilata/chat-service/infrastructure/grpc"
	"github.com/vulpes-ferrilata/chat-service/infrastructure/grpc/interceptors"
	"github.com/vulpes-ferrilata/chat-service/infrastructure/view/cassandra/projectors"
	"github.com/vulpes-ferrilata/chat-service/presentation/v1/servers"
	"go.uber.org/dig"
)

func NewContainer() *dig.Container {
	container := dig.New()

	//Infrastructure layer
	container.Provide(NewConfig)
	container.Provide(NewCassandra)
	container.Provide(NewValidator)
	container.Provide(NewLogrus)
	container.Provide(NewUniversalTranslator)
	container.Provide(grpc.NewServer)
	//--Grpc interceptors
	container.Provide(interceptors.NewRecoverInterceptor)
	container.Provide(interceptors.NewErrorHandlerInterceptor)
	container.Provide(interceptors.NewLocaleInterceptor)

	//Domain layer
	//--Repositories
	container.Provide(repositories.NewMessageRepository)

	//View layer
	//--Projectors
	container.Provide(projectors.NewMessageProjector)

	//Application layer
	//--Queries
	container.Provide(query_handlers.NewFindMessagesByRoomIDHandler)
	//--Commands
	container.Provide(command_handlers.NewCreateMessageCommandHandler)

	//Presentation layer
	container.Provide(servers.NewChatServer)

	return container
}
