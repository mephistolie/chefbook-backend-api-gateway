package recipe

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe/dto/common_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe/dto/request_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe/dto/response_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
)

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
			PictureLink: upload.PictureLink,
			UploadLink:  upload.UploadLink,
			FormData:    response.NonNilStringMap(upload.FormData),
			MaxSize:     upload.MaxSize,
		}
	}

	response.Success(c, gin.H{"uploads": uploads})
}

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

	res, err := h.service.SetRecipePictures(c, &api.SetRecipePicturesRequest{
		RecipeId:     c.Param(ParamRecipeId),
		UserId:       payload.UserId.String(),
		Pictures:     common_body.NewPicturesRequest(body.Pictures),
		Subscription: payload.SubscriptionPlan,
		Version:      body.Version,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.SetRecipePictures{
		Pictures: common_body.NewPicturesResponse(res.Pictures),
		Version:  res.Version,
	})
}
