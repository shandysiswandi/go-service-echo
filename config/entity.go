package config

// Config is
type Config struct {
	App      *AppConfig
	Database *DatabaseConfig
	Library  *LibraryConfig
	External *ExternalConfig
}

// AppConfig is
type AppConfig struct {
	Env      string
	Port     string
	Name     string
	Timezone string
}

// DatabaseConfig is
type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

// LibraryConfig is
type LibraryConfig struct {
	JWT       *JWTConfig
	Redis     *RedisConfig
	SentryDSN string
}

// JWTConfig is
type JWTConfig struct {
	AccessSecret  []byte
	RefreshSecret []byte
}

// RedisConfig is
type RedisConfig struct {
	Addr     string
	Password string
	Database int
}

// ExternalConfig is
type ExternalConfig struct {
	JsonplaceholderURL string
}
