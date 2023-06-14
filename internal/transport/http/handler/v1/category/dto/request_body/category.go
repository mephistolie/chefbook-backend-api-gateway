package request_body

import "github.com/google/uuid"

type AddCategory struct {
	Id    *uuid.UUID `json:"categoryId,omitempty"`
	Name  string     `json:"name"`
	Emoji *string    `json:"emoji,omitempty"`
}

type UpdateCategory struct {
	Name  string  `json:"name"`
	Emoji *string `json:"emoji,omitempty"`
}
