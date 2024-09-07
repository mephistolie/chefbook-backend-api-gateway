package response_body

import (
	common "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
)

func newProfilesInfo(profilesInfo map[string]*api.RecipeProfileInfo) map[string]common.ProfileMinInfo {
	dto := make(map[string]common.ProfileMinInfo)
	for id, profileInfo := range profilesInfo {
		dto[id] = common.ProfileMinInfo{
			Name:   profileInfo.Name,
			Avatar: profileInfo.Avatar,
		}
	}
	return dto
}
