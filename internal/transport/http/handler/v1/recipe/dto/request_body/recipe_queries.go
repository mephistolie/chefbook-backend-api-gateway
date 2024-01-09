package request_body

import (
	"github.com/google/uuid"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type GetRecipes struct {
	PageSize *int32 `json:"pageSize,omitempty"`

	RecipeIds *[]string `json:"recipeIds,omitempty"`
	AuthorId  *string   `json:"authorId,omitempty"`

	Owned bool `json:"owned,omitempty"`
	Saved bool `json:"saved,omitempty"`

	Search *string `json:"search,omitempty"`

	Sorting               *string    `json:"sorting,omitempty"`
	LastRecipeId          *string    `json:"lastRecipeId,omitempty"`
	LastCreationTimestamp *time.Time `json:"lastCreationTimestamp,omitempty"`
	LastUpdateTimestamp   *time.Time `json:"lastUpdateTimestamp,omitempty"`
	LastRating            *float32   `json:"lastRating,omitempty"`
	LastVotes             *int32     `json:"lastVotes,omitempty"`
	LastTime              *int32     `json:"lastTime,omitempty"`
	LastCalories          *int32     `json:"lastCalories,omitempty"`

	MinRating *int32 `json:"minRating,omitempty"`
	MaxRating *int32 `json:"maxRating,omitempty"`

	MinTime     *int32 `json:"minTime,omitempty"`
	MaxTime     *int32 `json:"maxTime,omitempty"`
	MinServings *int32 `json:"minServings,omitempty"`
	MaxServings *int32 `json:"maxServings,omitempty"`
	MinCalories *int32 `json:"minCalories,omitempty"`
	MaxCalories *int32 `json:"maxCalories,omitempty"`

	RecipeLanguages *[]string `json:"recipeLanguages,omitempty"`
	UserLanguage    *string   `json:"userLanguage,omitempty"`
}

func GetRecipesRequest(query GetRecipes, userId uuid.UUID) *api.GetRecipesRequest {
	var creationTimestamp *timestamppb.Timestamp
	if query.LastCreationTimestamp != nil {
		creationTimestamp = timestamppb.New(*query.LastCreationTimestamp)
	}
	var updateTimestamp *timestamppb.Timestamp
	if query.LastUpdateTimestamp != nil {
		updateTimestamp = timestamppb.New(*query.LastUpdateTimestamp)
	}
	var recipeLanguages []string
	if query.RecipeLanguages != nil {
		recipeLanguages = *query.RecipeLanguages
	}

	return &api.GetRecipesRequest{
		UserId:                userId.String(),
		PageSize:              query.PageSize,
		AuthorId:              query.AuthorId,
		Owned:                 query.Owned,
		Saved:                 query.Saved,
		Search:                query.Search,
		Sorting:               query.Sorting,
		LastRecipeId:          query.LastRecipeId,
		LastCreationTimestamp: creationTimestamp,
		LastUpdateTimestamp:   updateTimestamp,
		LastRating:            query.LastRating,
		LastVotes:             query.LastVotes,
		LastTime:              query.LastTime,
		LastCalories:          query.LastCalories,
		MinRating:             query.MinRating,
		MaxRating:             query.MaxRating,
		MinServings:           query.MinServings,
		MaxServings:           query.MaxServings,
		MinTime:               query.MinTime,
		MaxTime:               query.MaxTime,
		MinCalories:           query.MinCalories,
		MaxCalories:           query.MaxCalories,
		RecipeLanguages:       recipeLanguages,
		UserLanguage:          query.UserLanguage,
	}
}
