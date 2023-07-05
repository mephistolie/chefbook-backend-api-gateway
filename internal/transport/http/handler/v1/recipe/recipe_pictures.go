package recipe

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe/dto/request_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe/dto/response_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
)

// GenerateRecipePicturesUploadLinks Swagger Documentation
//
//	@Summary		Generate recipe pictures upload links
//	@Description	Generate recipe pictures upload links
//	@Tags			recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			recipe_id				path		string	true	"Recipe ID"
//	@Param			input		body		request_body.GenerateRecipePicturesUploadLinks	true	"Input"
//	@Success		200			{object}	[]response_body.RecipePictureUpload
//	@Failure		400			{object}	fail.Response
//	@Failure		500			{object}	fail.Response
//	@Router			/v1/recipes/{recipe_id}/pictures [post]
func (h *Handler) GenerateRecipePicturesUploadLinks(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.GenerateRecipePicturesUploadLinks
	if err = c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.service.GenerateRecipePicturesUploadLinks(c, &api.GenerateRecipePicturesUploadLinksRequest{
		RecipeId:      c.Param(ParamRecipeId),
		UserId:        payload.UserId.String(),
		PicturesCount: body.PicturesCount,
		Subscription:  payload.SubscriptionPlan,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	uploads := make([]response_body.RecipePictureUpload, len(res.Links))
	for i, upload := range res.Links {
		uploads[i] = response_body.RecipePictureUpload{
			PictureId: upload.PictureId,
			Link:      upload.Link,
			FormData:  upload.FormData,
			MaxSize:   upload.MaxSize,
		}
	}

	response.Success(c, uploads)
}

// SetRecipePictures Swagger Documentation
//
//	@Summary		Set recipe pictures
//	@Description	Set recipe pictures
//	@Tags			recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			recipe_id				path		string	true	"Recipe ID"
//	@Param			input		body		request_body.SetRecipePictures	true	"Pictures"
//	@Success		200			{object}	response_body.SetRecipePictures
//	@Failure		400			{object}	fail.Response
//	@Failure		500			{object}	fail.Response
//	@Router			/v1/recipes/{recipe_id}/pictures [post]
func (h *Handler) SetRecipePictures(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.SetRecipePictures
	if err = c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	var previewPtr *string
	if body.Preview != nil {
		preview := body.Preview.String()
		previewPtr = &preview
	}

	cooking := make(map[string]*api.StepPictures)
	if body.Cooking != nil {
		for step, pictures := range *body.Cooking {
			var rawPictures []string
			for _, picture := range pictures {
				rawPictures = append(rawPictures, picture.String())
			}
			cooking[step.String()] = &api.StepPictures{Pictures: rawPictures}
		}
	}

	res, err := h.service.SetRecipePictures(c, &api.SetRecipePicturesRequest{
		RecipeId:           c.Param(ParamRecipeId),
		UserId:             payload.UserId.String(),
		Preview:            previewPtr,
		CookingPicturesMap: cooking,
		Subscription:       payload.SubscriptionPlan,
		Version:            body.Version,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.SetRecipePictures{Version: res.Version})
}
