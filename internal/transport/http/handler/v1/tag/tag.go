package tag

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/tag/dto/response_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-tag/api/proto/implementation/v1"
)

const (
	queryLanguage = "language"
	queryGroups   = "group"
)

func (h *Handler) GetTags(c *gin.Context) {
	language, _ := c.GetQuery(queryLanguage)
	groups, _ := c.GetQueryArray(queryGroups)

	res, err := h.service.GetTags(c, &api.GetTagsRequest{LanguageCode: language, Groups: groups})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.GetTags(res))
}

func (h *Handler) GetTag(c *gin.Context) {
	language, _ := c.GetQuery(queryLanguage)

	res, err := h.service.GetTag(c, &api.GetTagRequest{
		TagId:        c.Param(ParamTagId),
		LanguageCode: language,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	var groupName *string
	if len(res.GroupName) > 0 {
		groupName = &res.GroupName
	}

	response.Success(c, response_body.GetTag(res.Tag, groupName))
}

func (h *Handler) GetTagGroups(c *gin.Context) {
	language, _ := c.GetQuery(queryLanguage)

	res, err := h.service.GetTagGroups(c, &api.GetTagGroupsRequest{LanguageCode: language})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, res.Groups)
}
