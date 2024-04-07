package model

type (
	// User model info
	// @Description User account information
	// @Description with user id and username
	UserDto struct {
		// ID this is userid
		ID string `json:"id"`
		// ID this is username
		Username string `json:"username"`
		Password string `json:"-"`
	}
)
