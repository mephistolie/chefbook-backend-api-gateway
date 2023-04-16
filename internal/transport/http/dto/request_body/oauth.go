package request_body

type OAuthCode struct {
	Code  string `json:"code"`
	State string `json:"state"`
}
