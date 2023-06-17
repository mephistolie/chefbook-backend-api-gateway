package request_body

type DeleteProfile struct {
	Password       string `json:"password"`
	WithSharedData bool   `json:"withSharedData,omitempty"`
}
