package config

// Config is
type Config struct {
	App      *AppConfig
	SSL      *SSLConfig
	Database *DatabaseConfig
	JWT      *JWTConfig
	Redis    *RedisConfig
	Sentry   *SentryConfig
	External *ExternalConfig
}

// AppConfig is
type AppConfig struct {
	Env      string
	Host     string
	Port     string
	Name     string
	Timezone string
}

// SSLConfig is
type SSLConfig struct {
	Cert string
	Key  string
}

// DatabaseConfig is
type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     string
	Username string
	Password string
	Name     string
	Timezone string
}

// JWTConfig is
type JWTConfig struct {
	AccessSecret  []byte
	RefreshSecret []byte
}

// RedisConfig is
type RedisConfig struct {
	Host     string
	Password string
	Database int
}

// SentryConfig is
type SentryConfig struct {
	DNS string
	ENV string
}

// ExternalConfig is
type ExternalConfig struct {
	JSONPlaceHolder string
}
