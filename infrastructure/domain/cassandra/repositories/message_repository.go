package repositories

import (
	"context"

	"github.com/pkg/errors"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/table"
	"github.com/vulpes-ferrilata/chat-service/domain/models"
	"github.com/vulpes-ferrilata/chat-service/domain/repositories"
	"github.com/vulpes-ferrilata/chat-service/infrastructure/domain/cassandra/mappers"
)

func NewMessageRepository(scyllaSession gocqlx.Session) repositories.MessageRepository {
	messageMetadata := table.Metadata{
		Name:    "message",
		Columns: []string{"id", "room_id", "user_id", "detail"},
		PartKey: []string{"room_id"},
		SortKey: []string{"id"},
	}

	messageTable := table.New(messageMetadata)

	return &messageRepository{
		scyllaSession: scyllaSession,
		messageTable:  messageTable,
	}
}

type messageRepository struct {
	scyllaSession gocqlx.Session
	messageTable  *table.Table
}

func (m messageRepository) Insert(ctx context.Context, message *models.Message) error {
	messageEntity := mappers.ToMessageEntity(message)

	if err := m.messageTable.InsertQueryContext(ctx, m.scyllaSession).BindStruct(messageEntity).ExecRelease(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
