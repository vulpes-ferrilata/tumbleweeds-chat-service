package models

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/vulpes-ferrilata/chat-service/infrastructure/cassandra/udt"
)

type MessageBuilder struct {
	id        gocql.UUID
	roomID    udt.ObjectID
	userID    udt.ObjectID
	detail    string
	createdAt time.Time
}

func (m MessageBuilder) SetID(id gocql.UUID) MessageBuilder {
	m.id = id

	return m
}

func (m MessageBuilder) SetRoomID(roomID udt.ObjectID) MessageBuilder {
	m.roomID = roomID

	return m
}

func (m MessageBuilder) SetUserID(userID udt.ObjectID) MessageBuilder {
	m.userID = userID

	return m
}

func (m MessageBuilder) SetDetail(detail string) MessageBuilder {
	m.detail = detail

	return m
}

func (m MessageBuilder) Create() *Message {
	return &Message{
		aggregateRoot: aggregateRoot{
			aggregate: aggregate{
				id: m.id,
			},
		},
		roomID: m.roomID,
		userID: m.userID,
		detail: m.detail,
	}
}
