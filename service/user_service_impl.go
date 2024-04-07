package service

import (
	"context"
	"golang-clean-architecture/dto/model"
	"golang-clean-architecture/dto/request"
	"golang-clean-architecture/dto/response"
	xerrors "golang-clean-architecture/errors"
	"golang-clean-architecture/repository"
	"golang-clean-architecture/target/ent"
	"net/http"

	"github.com/cetnfurkan/core/cache"
	"github.com/cetnfurkan/core/config"
	"github.com/cetnfurkan/core/errors"
	"github.com/cetnfurkan/core/mapper"
	"github.com/labstack/echo/v4"
)

type (
	UserServiceImpl struct {
		cache      cache.Cache
		cfg        *config.Server
		repository repository.UserRepository
	}
)

// NewUserServiceImpl creates a new user service instance.
//
// It takes a config instance, a user repository instance and a cache instance
// and returns a new user service interface instance.
func NewUserServiceImpl(cfg *config.Server, repository repository.UserRepository, cache cache.Cache) UserService {
	return &UserServiceImpl{
		cache:      cache,
		cfg:        cfg,
		repository: repository,
	}
}

func (service *UserServiceImpl) ListUsers(c echo.Context, dto *request.ListUsersDto) (*response.ListUsersDto, error) {
	ctx, cancel := context.WithTimeout(c.Request().Context(), service.cfg.RequestTimeout)
	defer cancel()

	users, err := service.repository.ListUsers(ctx)
	if err != nil {
		return nil, errors.HttpError(c, errors.HttpErrorInfo(
			xerrors.DATABASE_ERROR.Code, xerrors.DATABASE_ERROR.Message, http.StatusInternalServerError,
		))
	}

	usersDto, err := mapper.ToDto[[]*ent.User, []*model.UserDto](users)
	if err != nil {
		return nil, errors.HttpError(c, errors.HttpErrorInfo(
			xerrors.GENERIC_ERROR.Code, xerrors.GENERIC_ERROR.Message, http.StatusInternalServerError,
		))
	}

	// Example of cache usage
	/*
		data, _ := json.Marshal(usersDto)
		service.cache.Set("users", string(data), 0)
	*/

	// Example of redis client usage
	/*
		redisCache := service.cache.(*cache.RedisCache)
		redisCache.Client().HSet("users", "users", usersDto)
	*/

	// Example of rabbitmq producer and consumer usage
	/*
		rabbitmq := mq.NewRabbitMQ(&service.cfg.RabbitMQ)
		data, _ := json.Marshal(usersDto)
		rabbitmq.Producer().Produce("users", data)
		go rabbitmq.Consumer().Consume("users")
	*/

	return &response.ListUsersDto{
		Users: usersDto,
	}, nil
}
