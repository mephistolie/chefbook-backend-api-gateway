package response_body

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/contract"
	api "github.com/mephistolie/chefbook-backend-auth/api/proto/implementation/v1"
)

type Tokens = contract.TokensResponse

func NewTokens(res *api.SessionResponse) Tokens {
	restrictions := []contract.AccountRestriction{}
	if res.ProfileDeletionTimestamp != nil {
		var restriction contract.AccountRestriction
		_ = restriction.FromPendingDeletionRestriction(contract.PendingDeletionRestriction{Type: "pendingDeletion", DeletionTimestamp: res.ProfileDeletionTimestamp.AsTime(), DeleteSharedData: res.DeleteSharedData})
		restrictions = append(restrictions, restriction)
	}
	return Tokens{UserId: res.ProfileId, SessionId: res.SessionId, AccessToken: res.AccessToken, RefreshToken: res.RefreshToken, ExpirationTimestamp: res.ExpirationTimestamp.AsTime(), Restrictions: restrictions}
}
