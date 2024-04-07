package user

import (
	"net/http"

	"golang-clean-architecture/dto/request"
	xerrors "golang-clean-architecture/errors"

	"github.com/cetnfurkan/core/errors"
	"github.com/gorilla/schema"
	"github.com/labstack/echo/v4"
)

// Example of a controller method

// ListUsers godoc
// @Summary      Lists users
// @Description  List users registered in the system
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  response.ListUsersDto
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /users [get]

// ListUsers 	 godoc
// @Summary      Lists users
// @Description  List users registered in the system
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.ListUsersDto
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /users [get]
// @Security	 BearerAuth
func (userController *userController) listUsers(c echo.Context) error {
	var (
		decoder = schema.NewDecoder()
		reqDto  = new(request.ListUsersDto)
	)

	err := decoder.Decode(reqDto, c.QueryParams())
	if err != nil {
		return errors.HttpError(c, errors.HttpErrorInfo(
			xerrors.INVALID_PARAMETERS.Code, xerrors.INVALID_PARAMETERS.Message, http.StatusInternalServerError,
		))
	}
	/*
		// grpc client example
		client := userGrpcClient.NewUserClientImpl(userController.cfg)
		response, err := client.ListUsers(ctx, &grpc.ListUsersRequest{Query: ""})
		if err != nil {
			return errors.HttpError(c, errors.HttpErrorInfo(
				xerrors.GRPC_ERROR.Code, xerrors.GRPC_ERROR.Message, http.StatusInternalServerError,
			))
		}

		return c.JSON(http.StatusOK, response.Users)
	*/

	respDto, err := userController.service.ListUsers(c, reqDto)
	if respDto == nil || err != nil {
		return nil
	}

	return c.JSON(http.StatusOK, respDto)
}
