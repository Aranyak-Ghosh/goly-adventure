package main

import (
	"github.com/Aranyak-Ghosh/spotigo/user_profile/controllers"
	"github.com/Aranyak-Ghosh/spotigo/user_profile/models/database/user"
	"github.com/Aranyak-Ghosh/spotigo/user_profile/providers/config"
	"github.com/Aranyak-Ghosh/spotigo/user_profile/providers/database"
	"github.com/Aranyak-Ghosh/spotigo/user_profile/router"
	userService "github.com/Aranyak-Ghosh/spotigo/user_profile/services/user"
	"github.com/Aranyak-Ghosh/spotigo/user_profile/utils/http"
	"github.com/Aranyak-Ghosh/spotigo/user_profile/utils/logger"
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
