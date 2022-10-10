package infrastructure

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/vulpes-ferrilata/chat-service/infrastructure/config"
)

func NewCassandra(config config.Config) (gocqlx.Session, error) {
	cluster := gocql.NewCluster(config.Cassandra.Address)

	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: config.Cassandra.Username,
		Password: config.Cassandra.Password,
	}

	cluster.Keyspace = config.Cassandra.Keyspace

	return gocqlx.WrapSession(cluster.CreateSession())
}
