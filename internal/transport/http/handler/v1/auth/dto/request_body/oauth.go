package request_body

type GoogleOAuth struct {
	IdToken *string `json:"idToken"`
	Code    *string `json:"code"`
	State   *string `json:"state"`
}

type OAuthCode struct {
	Code  string `json:"code"`
	State string `json:"state"`
}
