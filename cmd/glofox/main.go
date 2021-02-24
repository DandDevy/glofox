package main

import (
	"github.com/DandDevy/glofox/api"
	"github.com/DandDevy/glofox/internal"
	"github.com/DandDevy/glofox/internal/repository"
	"github.com/labstack/echo/v4"
)

func main() {
	// Setup echo.
	e := echo.New()
	e.HideBanner = true
	//e.HidePort = true

	// Init db
	repoProvider := repository.NewRepositoryProvider(&repository.DB{})

	// Setup internal services
	serviceProvider := internal.NewServiceProvider(repoProvider)

	// Init Routes
	api.InitRoutes(e, serviceProvider)

	// Run echo server
	e.Logger.Fatal(e.Start(":8080"))
}
