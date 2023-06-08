package response_body

import (
	api "github.com/mephistolie/chefbook-backend-profile/api/proto/implementation/v1"
	"time"
)

type Profile struct {
	Id                    *string    `json:"profileId,omitempty"`
	Nickname              *string    `json:"nickname,omitempty"`
	Email                 *string    `json:"email,omitempty"`
	Role                  *string    `json:"role,omitempty"`
	OAuth                 *OAuth     `json:"oAuth,omitempty"`
	IsBlocked             bool       `json:"blocked"`
	RegistrationTimestamp *time.Time `json:"registeredAt,omitempty"`
	FirstName             *string    `json:"firstName,omitempty"`
	LastName              *string    `json:"lastName,omitempty"`
	Description           *string    `json:"description,omitempty"`
	Avatar                *string    `json:"avatar,omitempty"`
}

type OAuth struct {
	GoogleId *string `json:"googleId,omitempty"`
	VkId     *int64  `json:"vkId,omitempty"`
}

func GetProfile(profile *api.GetProfileResponse) Profile {
	res := Profile{
		IsBlocked: profile.IsBlocked,
	}
	if len(profile.Id) > 0 {
		res.Id = &profile.Id
	}
	if len(profile.Email) > 0 {
		res.Email = &profile.Email
	}
	if len(profile.Nickname) > 0 {
		res.Nickname = &profile.Nickname
	}
	if len(profile.Role) > 0 {
		res.Role = &profile.Role
	}
	if profile.OAuth != nil {
		oAuth := OAuth{}
		if len(profile.OAuth.GoogleId) > 0 {
			oAuth.GoogleId = &profile.OAuth.GoogleId
		}
		if profile.OAuth.VkId > 0 {
			oAuth.VkId = &profile.OAuth.VkId
		}
		res.OAuth = &oAuth
	}
	if profile.RegistrationTimestamp != nil {
		t := profile.RegistrationTimestamp.AsTime()
		res.RegistrationTimestamp = &t
	}

	if len(profile.FirstName) > 0 {
		res.FirstName = &profile.FirstName
	}
	if len(profile.LastName) > 0 {
		res.LastName = &profile.LastName
	}
	if len(profile.Description) > 0 {
		res.Description = &profile.Description
	}
	if len(profile.Avatar) > 0 {
		res.Avatar = &profile.Avatar
	}

	return res
}
