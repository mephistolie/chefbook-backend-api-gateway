package response_body

import (
	api "github.com/mephistolie/chefbook-backend-profile/api/proto/implementation/v1"
	"time"
)

type Profile struct {
	Id                    string     `json:"profileId,omitempty"`
	Nickname              *string    `json:"nickname,omitempty"`
	Email                 *string    `json:"email,omitempty"`
	Role                  *string    `json:"role,omitempty"`
	OAuth                 *OAuth     `json:"oAuth,omitempty"`
	IsBlocked             bool       `json:"blocked"`
	RegistrationTimestamp *time.Time `json:"registrationTimestamp,omitempty"`
	FirstName             *string    `json:"firstName,omitempty"`
	LastName              *string    `json:"lastName,omitempty"`
	Description           *string    `json:"description,omitempty"`
	Avatar                *string    `json:"avatar,omitempty"`
	SubscriptionPlan      string     `json:"subscriptionPlan,omitempty"`
}

type OAuth struct {
	GoogleId *string `json:"googleId,omitempty"`
	VkId     *int64  `json:"vkId,omitempty"`
}

func GetProfile(profile *api.GetProfileResponse) Profile {
	res := Profile{
		Id:               profile.Id,
		Nickname:         profile.Nickname,
		Email:            profile.Email,
		Role:             profile.Role,
		IsBlocked:        profile.IsBlocked,
		FirstName:        profile.FirstName,
		LastName:         profile.LastName,
		Description:      profile.Description,
		Avatar:           profile.Avatar,
		SubscriptionPlan: profile.SubscriptionPlan,
	}
	if profile.OAuth != nil {
		res.OAuth = &OAuth{
			GoogleId: profile.OAuth.GoogleId,
			VkId:     profile.OAuth.VkId,
		}
	}
	if profile.RegistrationTimestamp != nil {
		t := profile.RegistrationTimestamp.AsTime()
		res.RegistrationTimestamp = &t
	}

	return res
}
