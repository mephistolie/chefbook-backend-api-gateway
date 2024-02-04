package request_body

import "github.com/google/uuid"

type RequestPasswordReset struct {
	Email    string `json:"email,omitempty"`
	Nickname string `json:"nickname,omitempty"`
}

type ResetPassword struct {
	Id          uuid.UUID `json:"userId" binding:"required"`
	ResetCode   string    `json:"resetCode" binding:"required"`
	NewPassword string    `json:"newPassword" binding:"required"`
}

type ChangePassword struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
