package repositories

import (
	"context"

	"github.com/vulpes-ferrilata/chat-service/domain/models"
)

type MessageRepository interface {
	Insert(ctx context.Context, message *models.Message) error
}
