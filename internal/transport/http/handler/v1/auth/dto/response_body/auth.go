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
	ProfileId         string     `json:"profileId"`
	Access            string     `json:"accessToken"`
	Refresh           string     `json:"refreshToken"`
	ExpiresAt         time.Time  `json:"expirationTimestamp"`
	ProfileDeletingAt *time.Time `json:"profileDeletionTimestamp,omitempty"`
}

func NewTokens(res *api.SessionResponse) Tokens {
	var profileDeletionTimestamp *time.Time
	if res.ProfileDeletionTimestamp != nil {
		timestamp := res.ProfileDeletionTimestamp.AsTime()
		profileDeletionTimestamp = &timestamp
	}

	return Tokens{
		ProfileId:         res.ProfileId,
		Access:            res.AccessToken,
		Refresh:           res.RefreshToken,
		ExpiresAt:         res.ExpirationTimestamp.AsTime(),
		ProfileDeletingAt: profileDeletionTimestamp,
	}
}
