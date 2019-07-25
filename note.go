package gitlab

type Note struct {
	Id           int     `json:"id"`
	Note         string  `json:"note"`
	NoteableType string  `json:"noteable_type"`
	AuthorID     int     `json:"author_id"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
	ProjectID    int     `json:"project_id"`
	Attachment   *string `json:"attachment"`
	LineCode     *string `json:"line_code"`
	CommitID     string  `json:"commit_id"`
	NoteableID   *int    `json:"noteable_id"`
	System       bool    `json:"system"`
	StDiff       *Diff   `json:"st_diff"`
	Url          string  `json:"url"`
}
