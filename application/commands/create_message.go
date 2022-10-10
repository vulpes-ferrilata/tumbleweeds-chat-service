package commands

type CreateMessage struct {
	MessageID string `validate:"required,uuid"`
	RoomID    string `validate:"required,objectid"`
	UserID    string `validate:"required,objectid"`
	Detail    string `validate:"required"`
}
