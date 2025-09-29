package server

var Config struct {
	Server struct {
		Port string `env:"PORT" envDefault:"8080"`
	} `envprefix:"SERVER_"`
	Postgres struct {
		DBName   string `env:"DB_NAME" envDefault:"postgres"`
		User     string `env:"USER" envDefault:"postgres"`
		Password string `env:"PASSWORD" envDefault:"postgres"`
		Host     string `env:"HOST" envDefault:"localhost"`
		Port     string `env:"PORT" envDefault:"5432"`
	} `envprefix:"POSTGRES_"`
	Cassandra struct {
		Host string `env:"HOST" envDefault:"localhost"`
		Port string `env:"PORT" envDefault:"9042"`
	} `envprefix:"CASSANDRA_"`
}
