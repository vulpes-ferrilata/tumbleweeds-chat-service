package config

type Config struct {
	Server    ServerConfig    `mapstructure:"server"`
	Cassandra CassandraConfig `mapstructure:"cassandra"`
}
