package server

import (
	"golang-clean-architecture/config"
	"golang-clean-architecture/controller"
	"golang-clean-architecture/target/ent"

	"github.com/cetnfurkan/core/cache"
	"github.com/cetnfurkan/core/server"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewEchoServer creates a new echo server instance.
//
// It takes a config instance, a database instance and a cache instance
// and returns a new server interface instance.
//
// It will panic if it fails to create a new echo server instance.
func NewEchoServer(cfg *config.Config, db *ent.Client, cache cache.Cache) server.Server {
	return server.NewEchoServer(
		&cfg.Echo,
		server.WithControllers(controllers(cfg, db, cache)),
		server.WithMiddlewares(middleware.Logger()),
		server.WithSwaggerController(),
	)
}

func controllers(cfg *config.Config, db *ent.Client, cache cache.Cache) func(app *echo.Echo) {
	return func(app *echo.Echo) {
		controller.Init(app, cfg, db, cache)
	}
}
