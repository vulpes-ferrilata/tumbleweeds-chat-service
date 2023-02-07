package infrastructure

import (
	"github.com/vulpes-ferrilata/chat-service/application/queries"
	"github.com/vulpes-ferrilata/chat-service/infrastructure/cqrs/middlewares"
	"github.com/vulpes-ferrilata/cqrs"
)

func NewQueryBus(validationMiddleware *middlewares.ValidationMiddleware,
	findMessagesByRoomIDHandler *queries.FindMessagesByRoomIDHandler) (*cqrs.QueryBus, error) {
	queryBus := &cqrs.QueryBus{}

	queryBus.Use(
		validationMiddleware.QueryHandlerMiddleware(),
	)

	queryBus.Register(&queries.FindMessagesByRoomIDQuery{}, cqrs.WrapQueryHandlerFunc(findMessagesByRoomIDHandler.Handle))

	return queryBus, nil
}
