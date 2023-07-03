package response_body

import (
	api "github.com/mephistolie/chefbook-backend-tag/api/proto/implementation/v1"
)

type Tag struct {
	Id      string  `json:"tagId"`
	Name    string  `json:"name"`
	Emoji   *string `json:"emoji,omitempty"`
	GroupId *string `json:"groupId,omitempty"`
}

type TagsAndGroups struct {
	Tags   []Tag             `json:"tags"`
	Groups map[string]string `json:"groups,omitempty"`
}

type TagWithGroupName struct {
	Tag
	GroupName *string `json:"groupName,omitempty"`
}

func GetTags(response *api.GetTagsResponse) TagsAndGroups {
	tags := make([]Tag, len(response.Tags))
	for i, tag := range response.Tags {
		tags[i] = newTag(tag)
	}
	return TagsAndGroups{
		Tags:   tags,
		Groups: response.GroupNames,
	}
}

func GetTag(tag *api.Tag, group *string) TagWithGroupName {
	dto := newTag(tag)

	return TagWithGroupName{
		Tag:       dto,
		GroupName: group,
	}
}

func newTag(tag *api.Tag) Tag {
	return Tag{
		Id:      tag.TagId,
		Name:    tag.Name,
		Emoji:   tag.Emoji,
		GroupId: tag.GroupId,
	}
}

type AddCategory struct {
	Id string `json:"categoryId" binding:"required"`
}
