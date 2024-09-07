package request_body

type AddCollection struct {
	Id         *string `json:"collectionId,omitempty"`
	Name       string  `json:"name" binding:"required"`
	Visibility string  `json:"visibility,omitempty"`
}

type UpdateCollection struct {
	Name       string `json:"name" binding:"required"`
	Visibility string `json:"visibility,omitempty"`
}

type SaveCollectionToRecipeBook struct {
	ContributorKey *string `json:"contributorKey,omitempty"`
}
