package udt

import (
	"github.com/gocql/gocql"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ObjectID struct {
	primitive.ObjectID
}

func (o ObjectID) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	return []byte(o.ObjectID.Hex()), nil
}

func (o *ObjectID) UnmarshalCQL(info gocql.TypeInfo, data []byte) error {
	objectID, err := primitive.ObjectIDFromHex(string(data))
	if err != nil {
		return errors.WithStack(err)
	}

	o.ObjectID = objectID

	return nil
}
