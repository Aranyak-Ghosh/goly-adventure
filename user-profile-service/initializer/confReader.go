package initializer

import "gopkg.in/ini.v1"

type DatabaseConfig struct {
	Driver   string
	Server   string
	User     string
	Password string
	Database string
	Port     int32
}

func InitializeConfigs() (DatabaseConfig, error) {
	cfg, err := ini.Load("conf/config.ini")

	if err != nil {
		panic(err)
	}

	databaseConfig := DatabaseConfig{}
	cfg.Section("Database").MapTo(&databaseConfig)

	return databaseConfig, nil
}
