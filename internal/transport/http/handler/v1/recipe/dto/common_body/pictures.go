package common_body

type RecipePictures struct {
	Preview *string             `json:"preview,omitempty"`
	Cooking map[string][]string `json:"cooking,omitempty"`
}
