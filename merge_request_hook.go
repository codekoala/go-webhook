package gitlab

// See https://gitlab.com/help/user/project/integrations/webhooks.md#merge-request-events
type MergeRequestHook struct {
	Hook

	User             User         `json:"user"`
	Project          Project      `json:"project"`
	Repository       Repository   `json:"repository"`
	ObjectAttributes MergeRequest `json:"object_attributes"`
	Labels           []Label      `json:"labels"`
	Changes          Changes      `json:"changes"`
}
