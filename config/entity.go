package config

// Config is
type Config struct {
	App      app
	Database database
	Library  library
	External external
}

type app struct {
	Env      string
	Port     string
	Name     string
	Timezone string
}

type database struct {
	Driver   string
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

type library struct {
	JWT       jwt
	Redis     redis
	SentryDSN string
}

type jwt struct {
	AccessSecret  []byte
	RefreshSecret []byte
}

type redis struct {
	Addr     string
	Password string
	Database int
}

type external struct {
	JsonplaceholderURL string
}
