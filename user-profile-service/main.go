package main

import (
	"github.com/Aranyak-Ghosh/spotify/initializer"
	user "github.com/Aranyak-Ghosh/spotify/models/user"
)

func main() {

	dbConfig, ex := initializer.InitializeConfigs()

	if ex != nil {
		panic(ex)
	}

	db, err := initializer.InitializeDbContext(dbConfig)

	if err != nil {
		panic(err)
	}

	userRepo := user.UserDAO{}

	userRepo.Initialize(db)

	err = userRepo.Migrate()

	if err != nil {
		panic(err)
	}

	println("Hello, World!")
}
