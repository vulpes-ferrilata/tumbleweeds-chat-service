package projectors

import (
	"context"

	"github.com/vulpes-ferrilata/chat-service/view/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageProjector interface {
	FindByRoomID(ctx context.Context, roomID primitive.ObjectID) ([]*models.Message, error)
}
