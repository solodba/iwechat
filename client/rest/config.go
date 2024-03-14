package rest

type Config struct {
	Username string
	Password string
	URL      string
}

func NewConfig() *Config {
	return &Config{
		URL: "http://127.0.0.1:8888/ichatgpt/api/v1",
	}
}
