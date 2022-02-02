package database

import (
	"fmt"

	"github.com/Aranyak-Ghosh/spotigo/user_profile/providers/config"
	"go.uber.org/fx"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func NewDatabaseContext(configs *config.ConfigContainer) *gorm.DB {
	var dbConfig = configs.GetDatabaseConfig()
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", dbConfig.Server, dbConfig.User, dbConfig.Password, dbConfig.Database)

	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

var Module = fx.Option(fx.Provide(NewDatabaseContext))
