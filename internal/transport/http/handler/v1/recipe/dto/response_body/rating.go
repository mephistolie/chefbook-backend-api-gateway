package response_body

type Rating struct {
	Index float32 `json:"index"`
	Score *int32  `json:"score,omitempty"`
	Votes int32   `json:"votes"`
}
