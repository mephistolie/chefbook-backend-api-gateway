package request_body

import (
	"github.com/google/uuid"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type GetRecipes struct {
	RecipesCount *int32 `form:"count,omitempty"`

	RecipeIds []string `form:"recipe_id,omitempty"`
	AuthorId  *string  `form:"author_id,omitempty"`

	Owned bool `form:"owned,omitempty"`
	Saved bool `form:"saved,omitempty"`

	Tags []string `form:"tags,omitempty"`

	Search *string `form:"search,omitempty"`

	Sorting               *string    `form:"sorting,omitempty"`
	LastRecipeId          *string    `form:"last_recipe_id,omitempty"`
	LastCreationTimestamp *time.Time `form:"last_creation_timestamp,omitempty"`
	LastUpdateTimestamp   *time.Time `form:"last_update_timestamp,omitempty"`
	LastRating            *float32   `form:"last_rating,omitempty"`
	LastVotes             *int32     `form:"last_votes,omitempty"`
	LastTime              *int32     `form:"last_time,omitempty"`
	LastCalories          *int32     `form:"last_calories,omitempty"`

	MinRating *int32 `form:"min_rating,omitempty"`
	MaxRating *int32 `form:"max_rating,omitempty"`

	MinTime     *int32 `form:"min_time,omitempty"`
	MaxTime     *int32 `form:"max_time,omitempty"`
	MinServings *int32 `form:"min_servings,omitempty"`
	MaxServings *int32 `form:"max_servings,omitempty"`
	MinCalories *int32 `form:"min_calories,omitempty"`
	MaxCalories *int32 `form:"max_calories,omitempty"`

	RecipeLanguages []string `form:"recipe_language,omitempty"`
	UserLanguage    *string  `form:"user_language,omitempty"`
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

	return &api.GetRecipesRequest{
		UserId:                userId.String(),
		RecipeIds:             query.RecipeIds,
		PageSize:              query.RecipesCount,
		AuthorId:              query.AuthorId,
		Owned:                 query.Owned,
		Saved:                 query.Saved,
		Tags:                  query.Tags,
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
		RecipeLanguages:       query.RecipeLanguages,
		UserLanguage:          query.UserLanguage,
	}
}
