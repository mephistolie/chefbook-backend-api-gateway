package request_body

type ConfirmGoogleSubscription struct {
	SubscriptionId string `json:"subscriptionId"`
	PurchaseToken  string `json:"purchaseToken"`
}
