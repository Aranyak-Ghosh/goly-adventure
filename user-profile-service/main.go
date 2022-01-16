package main

import (
	"fmt"

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
	} else {
		err = userRepo.SeedData()

		if err != nil {
			panic(err)
		} else {
			var users []user.User
			var count int64
			var exc error
			users, count, exc = userRepo.List("", 0, 10)

			fmt.Printf("users %s, %s\n", users[0].ID, users[1].ID)
			err = userRepo.Follow(&users[0], &users[1])
			if err != nil {
				panic(err)
			}
			print(users)
			println(count)
			if exc != nil {
				panic(exc)
			}
		}
	}

	println("Hello, World!")
}
