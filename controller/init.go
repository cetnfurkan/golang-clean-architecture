package controller

import (
	"golang-clean-architecture/config"
	"golang-clean-architecture/controller/user"
	"golang-clean-architecture/target/ent"

	"github.com/cetnfurkan/core/cache"
	"github.com/labstack/echo/v4"
)

// Init initializes all controllers.
//
// It takes an echo instance, a config instance, a database instance and a cache instance.
func Init(app *echo.Echo, cfg *config.Config, db *ent.Client, cache cache.Cache) {
	// Add controllers here
	user.UserController(app, cfg, db, cache).Init()
}
