package infrastructure

import (
	"github.com/vulpes-ferrilata/chat-service/application/commands"
	"github.com/vulpes-ferrilata/chat-service/infrastructure/cqrs/middlewares"
	"github.com/vulpes-ferrilata/cqrs"
)

func NewCommandBus(validationMiddleware *middlewares.ValidationMiddleware,
	transactionMiddleware *middlewares.TransactionMiddleware,
	createMessageCommandHandler *commands.CreateMessageCommandHandler) (*cqrs.CommandBus, error) {
	commandBus := &cqrs.CommandBus{}

	commandBus.Use(
		validationMiddleware.CommandHandlerMiddleware(),
		transactionMiddleware.CommandHandlerMiddleware(),
	)

	commandBus.Register(&commands.CreateMessageCommand{}, cqrs.WrapCommandHandlerFunc(createMessageCommandHandler.Handle))

	return commandBus, nil
}
