package config

type DatabaseConfig struct {
	Enabled       bool   `env:"SERVER_RPC_DISABLED" default:"true" json:"enabled" yaml:"enabled" toml:"enabled"`
	Name          string `env:"DB_NAME" default:"admin_on_rest_entities" json:"name" yaml:"name" toml:"name"`
	Adapter       string `env:"DB_ADAPTER" default:"sqlite" json:"adapter" yaml:"adapter" toml:"adapter"`
	DSN           string `env:"DB_DSN" default:"host=localhost user=ireflect password=1Reflect dbname=ireflect-dev sslmode=disable" json:"dsn" yaml:"dsn" toml:"dsn"`
	Host          string `env:"DB_HOST" default:"localhost" json:"host" yaml:"host" toml:"host"`
	Port          string `env:"DB_PORT" default:"3306" json:"port" yaml:"port" toml:"port"`
	User          string `env:"DB_USER" json:"user" yaml:"user" toml:"user"`
	Password      string `env:"DB_PASSWORD" json:"password" yaml:"password" toml:"password"`
	PrefixPath    string `env:"DB_PREFIX_PATH" default:"./" json:"prefix_path" yaml:"prefix_path" toml:"prefix_path"`
	MigrationsDir string `env:"DB_MIGRATION_DIR" default:"./shared/data/seeds" json:"migrations_dir" yaml:"migrations_dir" toml:"migrations_dir"`
	SSL           bool   `env:"DB_SECURED" default:"false" json:"secured" yaml:"secured" toml:"secured"`
	SingularTable bool   `env:"DB_SINGULAR_TABLE" default:"true" json:"singular_table" yaml:"singular_table" toml:"singular_table"`
	Debug         bool   `env:"DB_DEBUG" default:"false" json:"debug" yaml:"debug" toml:"debug"`
}

type SourceConfig struct {
	Debug      bool `json:"debug" yaml:"debug" toml:"debug"`
	Disabled   bool `json:"disabled" yaml:"disabled" toml:"disabled"`
	Connection ConnectionConfig
	Collection CollectionConfig
}

type ConnectionConfig struct {
	Name          string
	Adapter       string
	Host          string
	Port          string
	User          string
	Password      string
	PrefixPath    string
	MigrationsDir string
	SSL           bool
	SingularTable bool
	Debug         bool
	DSN           string
}

type CollectionConfig struct {
	Debug    bool `json:"debug" yaml:"debug" toml:"debug"`
	Disabled bool `json:"disabled" yaml:"disabled" toml:"disabled"`
	Table    string
}
