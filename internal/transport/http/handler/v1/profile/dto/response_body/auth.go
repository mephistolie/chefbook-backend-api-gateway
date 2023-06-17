package response_body

import "time"

type ProfileDeletionStatus struct {
	DeletionTimestamp *time.Time `json:"deletionTimestamp,omitempty"`
	Deleted           bool       `json:"deleted,omitempty"`
}

type DeleteProfile struct {
	DeletionTimestamp time.Time `json:"deletionTimestamp"`
}
