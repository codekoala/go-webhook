package gitlab

type MergeRequest struct {
	Id              int      `json:"id"`
	TargetBranch    string   `json:"target_branch"`
	SourceBranch    string   `json:"source_branch"`
	SourceProjectID int      `json:"source_project_id"`
	AuthorID        int      `json:"author_id"`
	AssigneeID      *int     `json:"assignee_id"`
	Title           string   `json:"title"`
	CreatedAt       string   `json:"created_at"`
	UpdatedAt       string   `json:"updated_at"`
	StCommits       []Commit `json:"st_commits"`
	StDiffs         []Diff   `json:"st_diffs"`
	MilestoneId     *int     `json:"milestone_id"`
	State           string   `json:"state"`
	MergeStatus     string   `json:"merge_status"`
	TargetProjectID int      `json:"target_project_id"`
	IId             int      `json:"iid"`
	Description     string   `json:"description"`
	Position        int      `json:"position"`
	Source          Project  `json:"source"`
	Target          Project  `json:"target"`
	LastCommit      Commit   `json:"last_commit"`
	WorkInProgress  bool     `json:"work_in_progress"`
	Url             string   `json:"url"`
	Action          string   `json:"action"`
	Assignee        User     `json:"assignee"`
}
