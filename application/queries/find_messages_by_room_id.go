package queries

type FindMessagesByRoomID struct {
	RoomID string `validate:"omitempty,objectid"`
}
