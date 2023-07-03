package request_body

type AddCategory struct {
	Id    *string `json:"categoryId,omitempty"`
	Name  string  `json:"name"`
	Emoji *string `json:"emoji,omitempty"`
}

type UpdateCategory struct {
	Name  string  `json:"name"`
	Emoji *string `json:"emoji,omitempty"`
}
