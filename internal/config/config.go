package config

type Config struct {
	ServerAddr string
	AssetsDir  string
}

func NewConfig() Config {
	return Config{
		ServerAddr: ":8000",
		AssetsDir:  "./web/public/assets",
	}
}
