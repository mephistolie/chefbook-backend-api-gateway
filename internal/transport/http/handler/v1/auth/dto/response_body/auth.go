package response_body

import "time"

type SignUp struct {
	Id        string `json:"userId"`
	Activated bool   `json:"activated"`
	Message   string `json:"message,omitempty"`
}

type Tokens struct {
	Access    string    `json:"accessToken"`
	Refresh   string    `json:"refreshToken"`
	ExpiresAt time.Time `json:"expiresAt"`
}
