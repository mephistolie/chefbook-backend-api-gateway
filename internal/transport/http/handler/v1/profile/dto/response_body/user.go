package response_body

type GenerateAvatarUploadLink struct {
	AvatarId string            `json:"avatarId"`
	Link     string            `json:"link"`
	FormData map[string]string `json:"formData"`
}
