package infrastructure

import (
	"github.com/vulpes-ferrilata/chat-service/application/commands"
	"github.com/vulpes-ferrilata/chat-service/application/queries"
	"github.com/vulpes-ferrilata/chat-service/infrastructure/cqrs/middlewares"
	"github.com/vulpes-ferrilata/chat-service/infrastructure/domain/cassandra/repositories"
	"github.com/vulpes-ferrilata/chat-service/infrastructure/grpc/interceptors"
	"github.com/vulpes-ferrilata/chat-service/infrastructure/view/cassandra/projectors"
	"github.com/vulpes-ferrilata/chat-service/presentation"
	v1 "github.com/vulpes-ferrilata/chat-service/presentation/v1"
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
	//--Grpc interceptors
	container.Provide(interceptors.NewRecoverInterceptor)
	container.Provide(interceptors.NewErrorHandlerInterceptor)
	container.Provide(interceptors.NewLocaleInterceptor)
	//--Cqrs middlewares
	container.Provide(middlewares.NewValidationMiddleware)

	//Domain layer
	//--Repositories
	container.Provide(repositories.NewMessageRepository)

	//View layer
	//--Projectors
	container.Provide(projectors.NewMessageProjector)

	//Application layer
	//--Queries
	container.Provide(queries.NewFindMessagesByRoomIDHandler)
	//--Commands
	container.Provide(commands.NewCreateMessageCommandHandler)

	//Presentation layer
	container.Provide(presentation.NewServer)
	container.Provide(v1.NewChatServer)

	return container
}
