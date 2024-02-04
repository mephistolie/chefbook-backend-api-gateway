package response_body

import (
	api "github.com/mephistolie/chefbook-backend-auth/api/proto/implementation/v1"
	"time"
)

type SignUp struct {
	Id        string `json:"userId"`
	Activated bool   `json:"activated"`
	Message   string `json:"message,omitempty"`
}

type Tokens struct {
	Access            string     `json:"accessToken"`
	Refresh           string     `json:"refreshToken"`
	ExpiresAt         time.Time  `json:"expiresAt"`
	ProfileDeletingAt *time.Time `json:"profileDeletingAt,omitempty"`
}

func NewTokens(res *api.SessionResponse) Tokens {
	var profileDeletionTimestamp *time.Time
	if res.ProfileDeletionTimestamp != nil {
		timestamp := res.ProfileDeletionTimestamp.AsTime()
		profileDeletionTimestamp = &timestamp
	}

	return Tokens{
		Access:            res.AccessToken,
		Refresh:           res.RefreshToken,
		ExpiresAt:         res.ExpirationTimestamp.AsTime(),
		ProfileDeletingAt: profileDeletionTimestamp,
	}
}
