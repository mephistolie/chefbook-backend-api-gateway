package response_body

type Profile struct {
	Id     string  `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
	Avatar *string `json:"avatar,omitempty"`
}

type ProfileInfo struct {
	Name   *string `json:"name,omitempty"`
	Avatar *string `json:"avatar,omitempty"`
}
