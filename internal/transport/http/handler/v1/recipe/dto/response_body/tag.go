package response_body

import api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"

type Tag struct {
	Name  string  `json:"name"`
	Emoji *string `json:"emoji"`
}

func newTags(response map[string]*api.RecipeTag) map[string]Tag {
	tags := make(map[string]Tag)
	for id, tag := range response {
		tags[id] = newTag(tag)
	}
	return tags
}

func newTag(response *api.RecipeTag) Tag {
	return Tag{
		Name:  response.Name,
		Emoji: response.Emoji,
	}
}
