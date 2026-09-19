package request_body

import "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/contract"

type SignUp struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type RefreshToken = contract.RefreshSessionTokensRequest
type SignIn struct {
	Method   string `json:"method"`
	Login    string `json:"login"`
	Password string `json:"password"`
	IdToken  string `json:"idToken"`
	Code     string `json:"code"`
	State    string `json:"state"`
}
