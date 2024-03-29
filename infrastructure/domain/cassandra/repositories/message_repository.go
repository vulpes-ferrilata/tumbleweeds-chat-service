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

func NewMessageRepository(cassandraSession gocqlx.Session) repositories.MessageRepository {
	messageMetadata := table.Metadata{
		Name:    "messages",
		Columns: []string{"id", "room_id", "user_id", "detail"},
		PartKey: []string{"room_id"},
		SortKey: []string{"id"},
	}

	messageTable := table.New(messageMetadata)

	return &messageRepository{
		cassandraSession: cassandraSession,
		messageTable:     messageTable,
	}
}

type messageRepository struct {
	cassandraSession gocqlx.Session
	messageTable     *table.Table
}

func (m messageRepository) Insert(ctx context.Context, message *models.Message) error {
	messageEntity, err := mappers.MessageMapper.ToEntity(message)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := m.messageTable.InsertQueryContext(ctx, m.cassandraSession).BindStruct(messageEntity).ExecRelease(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
