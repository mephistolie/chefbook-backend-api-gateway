package response_body

type RecipePictureUpload struct {
	PictureId string            `json:"pictureId"`
	Link      string            `json:"link"`
	FormData  map[string]string `json:"formData"`
	MaxSize   int64             `json:"maxSize"`
}

type SetRecipePictures struct {
	Version int32 `json:"version"`
}
