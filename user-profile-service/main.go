package main

import (
	"github.com/Aranyak-Ghosh/spotify/controllers"
	"github.com/Aranyak-Ghosh/spotify/models/database/user"
	"github.com/Aranyak-Ghosh/spotify/providers/config"
	"github.com/Aranyak-Ghosh/spotify/providers/database"
	"github.com/Aranyak-Ghosh/spotify/router"
	userService "github.com/Aranyak-Ghosh/spotify/services/user"
	"github.com/Aranyak-Ghosh/spotify/utils/http"
	"github.com/Aranyak-Ghosh/spotify/utils/logger"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		http.Module,
		logger.Module,
		config.Module,
		database.Module,
		user.Module,
		userService.Module,
		controllers.Module,
		router.Module,
	)

	app.Run()
}
