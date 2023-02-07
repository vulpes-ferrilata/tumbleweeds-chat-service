package config

type CassandraConfig struct {
	Address  string `mapstructure:"address"`
	Keyspace string `mapstructure:"keyspace"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
