package response_body

type Contributor struct {
	Id   string `json:"userId" binding:"required"`
	Role string `json:"role" binding:"required"`
}
