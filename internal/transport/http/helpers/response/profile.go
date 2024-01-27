package response

type ProfileInfo struct {
	Id     string  `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
	Avatar *string `json:"avatar,omitempty"`
}
