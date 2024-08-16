package configs

type Config struct {
	Database struct {
		URL string `toml:"url" validate:"required,url"`
	} `toml:"database"`
}
