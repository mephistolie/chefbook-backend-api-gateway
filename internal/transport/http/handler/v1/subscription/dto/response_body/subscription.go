package response_body

import (
	api "github.com/mephistolie/chefbook-backend-subscription/api/proto/implementation/v1"
	"time"
)

type Subscription struct {
	Plan                string     `json:"plan"`
	Source              *string    `json:"source,omitempty"`
	ExpirationTimestamp *time.Time `json:"expirationTimestamp,omitempty"`
	AutoRenew           bool       `json:"autoRenew,omitempty"`
}

func GetSubscriptions(response []*api.Subscription) []Subscription {
	subscriptions := make([]Subscription, len(response))
	for i, subscription := range response {
		var expirationDatePtr *time.Time = nil
		if subscription.ExpirationDate != nil {
			expirationDate := subscription.ExpirationDate.AsTime()
			expirationDatePtr = &expirationDate
		}

		subscriptions[i] = Subscription{
			Plan:                subscription.Plan,
			Source:              subscription.Source,
			ExpirationTimestamp: expirationDatePtr,
			AutoRenew:           subscription.AutoRenew,
		}
	}
	return subscriptions
}
