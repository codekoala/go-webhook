package gitlab

type Issue struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	AssigneeIDs []int   `json:"assignee_ids"`
	AssigneeID  *int    `json:"assignee_id"`
	AuthorID    int     `json:"author_id"`
	ProjectID   int     `json:"project_id"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	Position    int     `json:"position"`
	BranchName  *string `json:"branch_name"`
	Description string  `json:"description"`
	MilestoneId *int    `json:"milestone_id"`
	State       string  `json:"state"`
	IId         int     `json:"iid"`
	Url         string  `json:"url"`
	Action      string  `json:"action"`
}
