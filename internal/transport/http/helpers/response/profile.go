package response

type ProfileMinInfo struct {
	Name   *string `json:"name,omitempty"`
	Avatar *string `json:"avatar,omitempty"`
}

type ProfileInfo struct {
	Id     string  `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
	Avatar *string `json:"avatar,omitempty"`
}
