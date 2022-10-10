package models

import (
	"time"

	"github.com/gocql/gocql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID        gocql.UUID
	RoomID    primitive.ObjectID
	UserID    primitive.ObjectID
	Detail    string
	CreatedAt time.Time
}
