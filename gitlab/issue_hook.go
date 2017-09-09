package gitlab

type IssueHook struct {
	Hook

	User             User       `json:"user"`
	Project          Project    `json:"project"`
	Repository       Repository `json:"repository"`
	ObjectAttributes Issue      `json:"object_attributes"`
	Assignees        []User     `json:"assignees"`
	Assignee         *User      `json:"assignee"`
	Labels           []Label    `json:"labels"`
}
