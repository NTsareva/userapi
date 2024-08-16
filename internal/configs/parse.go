package configs

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ParseAndValidate(filename string) (Config, error) {
	var config Config

	if _, err := toml.DecodeFile(filename, &config); err != nil {
		return Config{}, fmt.Errorf("can't parse toml file: %v", err)
	}

	err := validate.Struct(config)
	if err != nil {
		return Config{}, fmt.Errorf("can't validate config file: %v", err)
	}

	return config, nil
}
