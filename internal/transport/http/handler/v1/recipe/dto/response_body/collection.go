package response_body

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
)

type Collection struct {
	Id           string        `json:"collectionId" binding:"required"`
	Name         string        `json:"name" binding:"required"`
	Visibility   string        `json:"visibility" binding:"required"`
	Contributors []Contributor `json:"contributors"`
	RecipesCount int32         `json:"recipesCount"`
}

type CollectionInfo struct {
	Name string `json:"name" binding:"required"`
}

type AddCollection struct {
	Id string `json:"collectionId" binding:"required"`
}

type GetCollections struct {
	Collections  []Collection                       `json:"collections" binding:"required"`
	ProfilesInfo map[string]response.ProfileMinInfo `json:"profilesInfo"`
}

type GetCollection struct {
	Collection   Collection                         `json:"collection" binding:"required"`
	ProfilesInfo map[string]response.ProfileMinInfo `json:"profilesInfo"`
}

func newCollections(response []*api.Collection) []Collection {
	dtos := make([]Collection, len(response))
	for i, collection := range response {
		dtos[i] = newCollection(collection)
	}
	return dtos
}

func newCollectionsMap(response map[string]*api.CollectionInfo) map[string]CollectionInfo {
	collections := make(map[string]CollectionInfo)
	for id, collection := range response {
		collections[id] = newCollectionInfo(collection)
	}
	return collections
}

func NewGetCollections(response *api.GetCollectionsResponse) GetCollections {
	return GetCollections{
		Collections:  newCollections(response.Collections),
		ProfilesInfo: newProfilesInfo(response.ProfilesInfo),
	}
}

func NewGetCollection(response *api.GetCollectionResponse) GetCollection {
	return GetCollection{
		Collection:   newCollection(response.Collection),
		ProfilesInfo: newProfilesInfo(response.ProfilesInfo),
	}
}

func newCollection(collection *api.Collection) Collection {
	contributors := make([]Contributor, len(collection.Contributors))
	for i, contributor := range collection.Contributors {
		contributors[i] = Contributor{
			Id:   contributor.ContributorId,
			Role: contributor.Role,
		}
	}
	return Collection{
		Id:           collection.CollectionId,
		Name:         collection.Name,
		Visibility:   collection.Visibility,
		Contributors: contributors,
		RecipesCount: collection.RecipesCount,
	}
}

func newCollectionInfo(response *api.CollectionInfo) CollectionInfo {
	return CollectionInfo{
		Name: response.Name,
	}
}
