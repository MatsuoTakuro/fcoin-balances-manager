package api

type ErrResponse struct {
	Message string   `json:"message"`
	Details []string `json:"details,omitempty"`
}
