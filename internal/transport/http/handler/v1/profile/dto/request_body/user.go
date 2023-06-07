package request_body

import "github.com/google/uuid"

type SetName struct {
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
}

type SetDescription struct {
	Description *string `json:"description,omitempty"`
}

type ConfirmAvatarUploading struct {
	AvatarId uuid.UUID `json:"avatarId"`
}
