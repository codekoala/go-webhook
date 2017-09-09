package gitlab

type NoteHook struct {
	Hook

	User             User         `json:"user"`
	ProjectID        int          `json:"project_id"`
	Project          Project      `json:"project"`
	Repository       Repository   `json:"repository"`
	ObjectAttributes Note         `json:"object_attributes"`
	Commit           Commit       `json:"commit"`
	MergeRequest     MergeRequest `json:"merge_request"`
	Issue            Issue        `json:"issue"`
	Snippet          Snippet      `json:"snippet"`
}
