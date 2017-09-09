package gitlab

type Commit struct {
	Id        string   `json:"id"`
	Message   string   `json:"message"`
	Timestamp string   `json:"timestamp"`
	Url       string   `json:"url"`
	Author    User     `json:"author,omitempty"`
	Added     []string `json:"added"`
	Modified  []string `json:"modified"`
	Removed   []string `json:"removed"`
}

type BuildCommit struct {
	Id          int     `json:"id"`
	Sha         string  `json:"sha"`
	Message     string  `json:"message"`
	AuthorName  string  `json:"author_name,omitempty"`
	AuthorEmail string  `json:"author_email,omitempty"`
	Status      string  `json:"status"`
	Duration    *int    `json:"duration"`
	StartedAt   *string `json:"started_at"`
	FinishedAt  *string `json:"finished_at"`
}
