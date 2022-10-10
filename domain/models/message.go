package models

import (
	"github.com/vulpes-ferrilata/chat-service/infrastructure/cassandra/udt"
)

type Message struct {
	aggregateRoot
	roomID udt.ObjectID
	userID udt.ObjectID
	detail string
}

func (m Message) GetRoomID() udt.ObjectID {
	return m.roomID
}

func (m Message) GetUserID() udt.ObjectID {
	return m.userID
}

func (m Message) GetDetail() string {
	return m.detail
}
