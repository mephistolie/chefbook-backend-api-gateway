package response_body

type GetRecipeKey struct {
	Key *string `json:"key"`
}

type RecipeKeyRequest struct {
	UserId     string  `json:"userId"`
	UserName   *string `json:"userName"`
	UserAvatar *string `json:"userAvatar"`
	Status     string  `json:"status"`
	PublicKey  *string `json:"publicKey"`
}
