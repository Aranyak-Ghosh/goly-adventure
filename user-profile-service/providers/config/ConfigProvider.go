package config

import (
	"os"

	"github.com/Aranyak-Ghosh/spotify/models/config"
	"go.uber.org/fx"
	"gopkg.in/ini.v1"
)

type ConfigContainer struct {
	appConfig config.ApplicationConfig
	dbConfig  config.DatabaseConfig
}

func (container *ConfigContainer) GetApplicationConfig() config.ApplicationConfig {
	return container.appConfig
}

func (container *ConfigContainer) GetDatabaseConfig() config.DatabaseConfig {
	return container.dbConfig
}

func NewConfigContainer() (*ConfigContainer, error) {

	env := os.Getenv("ENV")
	file := ""
	if env == "" {
		file = "conf/config.ini"
	} else {
		file = "conf/config." + env + ".ini"
	}

	cfg, err := ini.Load(file)

	if err != nil {
		panic(err)
	}

	var configContainer = ConfigContainer{}

	cfg.Section(config.DbKey).MapTo(&configContainer.dbConfig)
	cfg.Section(config.ApplicationKey).MapTo(&configContainer.appConfig)

	if err != nil {
		return nil, err
	}
	return &configContainer, nil
}

func GetEnvironment() string {
	return "dev"
}

var Module = fx.Option(fx.Provide(NewConfigContainer))
