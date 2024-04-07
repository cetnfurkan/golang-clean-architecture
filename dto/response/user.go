package response

import (
	"golang-clean-architecture/dto/model"
)

type (
	ListUsersDto struct {
		Users []*model.UserDto `json:"users"`
	}
)
