package service

import (
	"golang-clean-architecture/dto/request"
	"golang-clean-architecture/dto/response"

	"github.com/labstack/echo/v4"
)

type (
	UserService interface {
		// ListUsers returns a list of users.
		//
		// It takes a context and a list users dto request instance and
		// returns a list users dto response instance and an error.
		//
		// It will map user entities from repository to user dto and return it.
		ListUsers(echo.Context, *request.ListUsersDto) (*response.ListUsersDto, error)
	}
)
