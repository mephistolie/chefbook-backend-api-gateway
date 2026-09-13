package request_body

type SetDisplayName struct {
	DisplayName *string `json:"displayName,omitempty"`
}

type SetDescription struct {
	Description *string `json:"description,omitempty"`
}

type ConfirmAvatarUploading struct {
	AvatarLink string `json:"avatarLink" binding:"required"`
}
