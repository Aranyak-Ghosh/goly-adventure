package config

var DbKey = "Database"

type DatabaseConfig struct {
	Driver   string
	Server   string
	User     string
	Password string
	Database string
	Port     int32
}
