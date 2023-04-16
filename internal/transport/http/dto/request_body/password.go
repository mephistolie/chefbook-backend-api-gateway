package request_body

import "github.com/google/uuid"

type RequestPasswordReset struct {
	Email    string `json:"email,omitempty"`
	Nickname string `json:"nickname,omitempty"`
}

type ResetPassword struct {
	Id          uuid.UUID `json:"userId"`
	ResetCode   string    `json:"resetCode"`
	NewPassword string    `json:"newPassword"`
}

type ChangePassword struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
