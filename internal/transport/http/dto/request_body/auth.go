package request_body

import (
	"github.com/google/uuid"
)

type SignUp struct {
	Id       *uuid.UUID `json:"userId"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
}

type SignIn struct {
	Email    string `json:"email,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Password string `json:"password"`
}

type RefreshToken struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

type DeleteProfile struct {
	Password string `json:"password"`
}
