package request_body

type SetName struct {
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
}

type SetDescription struct {
	Description *string `json:"description,omitempty"`
}

type ConfirmAvatarUploading struct {
	AvatarLink string `json:"avatarLink" binding:"required"`
}
