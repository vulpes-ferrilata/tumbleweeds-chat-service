package entities

import (
	"github.com/gocql/gocql"
	"github.com/vulpes-ferrilata/chat-service/infrastructure/cassandra/udt"
)

type Message struct {
	ID     gocql.UUID   `db:"id"`
	RoomID udt.ObjectID `db:"room_id"`
	UserID udt.ObjectID `db:"user_id"`
	Detail string       `db:"detail"`
}
