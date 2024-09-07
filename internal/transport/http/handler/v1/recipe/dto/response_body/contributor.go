package response_body

type Contributor struct {
	Id   string `json:"contributorId" binding:"required"`
	Role string `json:"role" binding:"required"`
}
