package models

import "github.com/gocql/gocql"

type aggregate struct {
	id gocql.UUID
}

func (a aggregate) GetID() gocql.UUID {
	return a.id
}
