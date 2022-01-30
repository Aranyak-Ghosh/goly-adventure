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

	cfg, err := ini.Load("config/config." + os.Getenv("ENV") + ".ini")

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
