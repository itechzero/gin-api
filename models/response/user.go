package response

import (
	"gin-api/models"
)

type SysUserResponse struct {
	User models.User `json:"user"`
}

type LoginResponse struct {
	User      models.User `json:"user"`
	Token     string      `json:"token"`
	ExpiresAt int64       `json:"expiresAt"`
}
