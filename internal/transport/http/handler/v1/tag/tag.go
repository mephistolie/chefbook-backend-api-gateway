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

// GetTags Swagger Documentation
//
//	@Summary		Get tags
//	@Description	Get recipe tags
//	@Tags			tag, recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			language			query		string		false	"Language code"
//	@Param			group				query		[]string	false	"Tag groups"
//	@Success		200					{object}	response_body.TagsAndGroups
//	@Failure		400					{object}	fail.Response
//	@Failure		500					{object}	fail.Response
//	@Router			/v1/recipes/tags	[get]
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

// GetTag Swagger Documentation
//
//	@Summary		Get tag
//	@Description	Get recipe tag
//	@Tags			tag, recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			tag_id						path		string	true	"Tag ID"
//	@Param			language					query		string	false	"Language code"
//	@Success		200							{object}	response_body.TagWithGroupName
//	@Failure		400							{object}	fail.Response
//	@Failure		500							{object}	fail.Response
//	@Router			/v1/recipes/tags/{tag_id}	[get]
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

// GetTagGroups Swagger Documentation
//
//	@Summary		Get tag groups
//	@Description	Get recipe tag groups
//	@Tags			tag, recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			language				query		string	false	"Language code"
//	@Success		200						{object}	map[string]string
//	@Failure		400						{object}	fail.Response
//	@Failure		500						{object}	fail.Response
//	@Router			/v1/recipes/tags/groups	[get]
func (h *Handler) GetTagGroups(c *gin.Context) {
	language, _ := c.GetQuery(queryLanguage)

	res, err := h.service.GetTagGroups(c, &api.GetTagGroupsRequest{LanguageCode: language})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, res.Groups)
}
