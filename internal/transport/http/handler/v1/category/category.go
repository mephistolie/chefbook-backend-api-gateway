package category

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/category/dto/request_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/category/dto/response_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-category/api/proto/implementation/v1"
)

// GetCategories Swagger Documentation
//
//	@Summary		Get user categories
//	@Description	Get user recipe categories
//	@Tags			category, recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Success		200						{object}	[]response_body.Category
//	@Failure		400						{object}	fail.Response
//	@Failure		500						{object}	fail.Response
//	@Router			/v1/recipes/categories	[get]
func (h *Handler) GetCategories(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.GetUserCategories(c, &api.GetUserCategoriesRequest{UserId: payload.UserId.String()})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.GetCategories(res))
}

// AddCategory Swagger Documentation
//
//	@Summary		Add category
//	@Description	Add recipe category
//	@Tags			category, recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			input					body		request_body.AddCategory	true	"Input"
//	@Success		200						{object}	response_body.AddCategory
//	@Failure		400						{object}	fail.Response
//	@Failure		500						{object}	fail.Response
//	@Router			/v1/recipes/categories	[post]
func (h *Handler) AddCategory(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.AddCategory
	if err = c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.service.CreateCategory(c, &api.CreateCategoryRequest{
		UserId:     payload.UserId.String(),
		CategoryId: body.Id,
		Name:       body.Name,
		Emoji:      body.Emoji,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.AddCategory{Id: res.CategoryId})
}

// GetCategory Swagger Documentation
//
//	@Summary		Get category
//	@Description	Get recipe category
//	@Tags			category, recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			category_id								path		string	true	"Category ID"
//	@Success		200										{object}	response_body.Category
//	@Failure		400										{object}	fail.Response
//	@Failure		500										{object}	fail.Response
//	@Router			/v1/recipes/categories/{category_id}	[get]
func (h *Handler) GetCategory(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.GetCategory(c, &api.GetCategoryRequest{
		CategoryId: c.Param(ParamCategoryId),
		UserId:     payload.UserId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.Category{
		Id:    res.CategoryId,
		Name:  res.Name,
		Emoji: res.Emoji,
	})
}

// UpdateCategory Swagger Documentation
//
//	@Summary		Update category
//	@Description	Update recipe category
//	@Tags			category, recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			category_id								path		int							true	"Category ID"
//	@Param			input									body		request_body.UpdateCategory	true	"Input"
//	@Success		200										{object}	response.MessageBody
//	@Failure		400										{object}	fail.Response
//	@Failure		500										{object}	fail.Response
//	@Router			/v1/recipes/categories/{category_id}	[put]
func (h *Handler) UpdateCategory(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.UpdateCategory
	if err = c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.service.UpdateCategory(c, &api.UpdateCategoryRequest{
		UserId:     payload.UserId.String(),
		CategoryId: c.Param(ParamCategoryId),
		Name:       body.Name,
		Emoji:      body.Emoji,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

// DeleteCategory Swagger Documentation
//
//	@Summary		Delete category
//	@Description	Delete recipe category
//	@Tags			category
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			category_id								path		int	true	"Category ID"
//	@Success		200										{object}	response.MessageBody
//	@Failure		400										{object}	fail.Response
//	@Failure		500										{object}	fail.Response
//	@Router			/v1/recipes/categories/{category_id}	[delete]
func (h *Handler) DeleteCategory(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.DeleteCategory(c, &api.DeleteCategoryRequest{
		CategoryId: c.Param(ParamCategoryId),
		UserId:     payload.UserId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}
