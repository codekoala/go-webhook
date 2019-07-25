package gitlab

type Changes struct {
	UpdatedById []*int        `json:"updated_by_id"`
	UpdatedAt   []string      `json:"updated_at"`
	Labels      ChangesLabels `json:"labels"`
}

type ChangesLabels struct {
	Previous []Label `json:"previous"`
	Current  []Label `json:"current"`
}
