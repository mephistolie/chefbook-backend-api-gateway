package response_body

type GenerateAvatarUploadLink struct {
	PictureLink string            `json:"pictureLink"`
	UploadLink  string            `json:"uploadLink"`
	FormData    map[string]string `json:"formData"`
	MaxSize     int64             `json:"maxSize"`
}
