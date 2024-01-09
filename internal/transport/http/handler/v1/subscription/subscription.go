package subscription

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/subscription/dto/request_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/subscription/dto/response_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-subscription/api/proto/implementation/v1"
)

// GetSubscriptions Swagger Documentation
//
//	@Summary		Get subscriptions
//	@Description	Get subscriptions
//	@Tags			subscription
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Success		200					{object}	[]response_body.Subscription
//	@Failure		400					{object}	fail.Response
//	@Failure		500					{object}	fail.Response
//	@Router			/v1/subscriptions	[get]
func (h *Handler) GetSubscriptions(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.GetProfileSubscriptions(c, &api.GetProfileSubscriptionsRequest{UserId: payload.UserId.String()})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.GetSubscriptions(res.Subscriptions))
}

// ConfirmGoogleSubscription Swagger Documentation
//
//	@Summary		Confirm Google subscription
//	@Description	Confirm Google subscription
//	@Tags			subscription
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			input						body		request_body.ConfirmGoogleSubscription	true	"Purchase Token"
//	@Success		200							{object}	response.MessageBody
//	@Failure		400							{object}	fail.Response
//	@Failure		500							{object}	fail.Response
//	@Router			/v1/subscriptions/google	[post]
func (h *Handler) ConfirmGoogleSubscription(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.ConfirmGoogleSubscription
	if err = c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.service.ConfirmGoogleSubscription(c, &api.ConfirmGoogleSubscriptionRequest{
		UserId:         payload.UserId.String(),
		SubscriptionId: body.SubscriptionId,
		PurchaseToken:  body.PurchaseToken,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}
