package user

import (
	"golang-clean-architecture/config"
	"golang-clean-architecture/repository"
	"golang-clean-architecture/service"
	"golang-clean-architecture/target/ent"

	"github.com/cetnfurkan/core/cache"
	"github.com/labstack/echo/v4"
)

type (
	userController struct {
		app     *echo.Echo
		cfg     *config.Config
		service service.UserService
	}
)

func UserController(app *echo.Echo, cfg *config.Config, db *ent.Client, cache cache.Cache) *userController {
	userRepository := repository.NewUserPostgresRepository(db)
	userService := service.NewUserServiceImpl(&cfg.Echo, userRepository, cache)

	return &userController{
		app:     app,
		cfg:     cfg,
		service: userService,
	}
}

func (controller *userController) Init() {
	group := controller.app.Group("/users")

	group.GET("", controller.listUsers)
}
