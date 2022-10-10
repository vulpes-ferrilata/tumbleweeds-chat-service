package projectors

import (
	"context"

	"github.com/pkg/errors"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
	"github.com/vulpes-ferrilata/chat-service/infrastructure/utils/slices"
	"github.com/vulpes-ferrilata/chat-service/infrastructure/view/cassandra/entities"
	"github.com/vulpes-ferrilata/chat-service/infrastructure/view/cassandra/mappers"
	"github.com/vulpes-ferrilata/chat-service/view/models"
	"github.com/vulpes-ferrilata/chat-service/view/projectors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewMessageProjector(scyllaSession gocqlx.Session) projectors.MessageProjector {
	messageMetadata := table.Metadata{
		Name:    "message",
		Columns: []string{"id", "room_id", "user_id", "detail"},
		PartKey: []string{"room_id"},
		SortKey: []string{"id"},
	}

	messageTable := table.New(messageMetadata)

	return &messageProjector{
		scyllaSession: scyllaSession,
		messageTable:  messageTable,
	}
}

type messageProjector struct {
	scyllaSession gocqlx.Session
	messageTable  *table.Table
}

func (m messageProjector) FindByRoomID(ctx context.Context, roomID primitive.ObjectID) ([]*models.Message, error) {
	messageEntities := make([]*entities.Message, 0)

	if err := m.messageTable.SelectQueryContext(ctx, m.scyllaSession).BindMap(qb.M{"room_id": roomID.Hex()}).SelectRelease(&messageEntities); err != nil {
		return nil, errors.WithStack(err)
	}

	messages, err := slices.Map(func(messageEntity *entities.Message) (*models.Message, error) {
		return mappers.ToMessageView(messageEntity), nil
	}, messageEntities)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return messages, nil
}
