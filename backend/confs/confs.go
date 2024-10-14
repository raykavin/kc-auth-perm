package confs

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		Application *Application `yaml:"application"`
	}

	Application struct {
		LoggerLevel string `yaml:"logger_level"`
		Web         *Web   `yaml:"web"`
		OIDC        *OIDC  `yaml:"oidc"`
	}

	Web struct {
		Listen uint   `yaml:"listen"`
		Ssl    bool   `yaml:"ssl"`
		Crt    string `yaml:"crt"`
		Key    string `yaml:"key"`
	}

	OIDC struct {
		ConfigurationUrl string `yaml:"configuration_url"`
		ClientID         string `yaml:"client_id"`
	}
)

var Default = &Config{}

func Load(path string) error {
	if len(strings.Trim(path, " ")) == 0 {
		return fmt.Errorf("invalid yaml config")
	}

	f, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(f, &Default)
}
