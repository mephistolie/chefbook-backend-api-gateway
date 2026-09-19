package subscription

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/subscription/dto/request_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/subscription/dto/response_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-subscription/api/proto/implementation/v1"
)

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

	response.Success(c, gin.H{"subscriptions": response_body.GetSubscriptions(res.Subscriptions)})
}

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
