package gitlab

type Snippet struct {
	Id              int    `json:"id"`
	Title           string `json:"title"`
	Content         string `json:"content"`
	AuthorID        int    `json:"author_id"`
	ProjectID       int    `json:"project_id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	FileName        string `json:"file_name"`
	ExpiresAt       string `json:"expires_at"`
	Type            string `json:"type"`
	VisibilityLevel int    `json:"visibility_level"`
}
