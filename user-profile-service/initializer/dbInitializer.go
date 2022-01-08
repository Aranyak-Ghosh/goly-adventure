package initializer

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func InitializeDbContext(dbConfig DatabaseConfig) (*gorm.DB, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", dbConfig.Server, dbConfig.User, dbConfig.Password, dbConfig.Database)

	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db, err
}
