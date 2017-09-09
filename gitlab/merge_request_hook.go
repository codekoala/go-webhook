package gitlab

// See https://gitlab.com/help/user/project/integrations/webhooks.md#merge-request-events
type MergeRequestHook struct {
	Hook

	User             User         `json:"user"`
	ObjectAttributes MergeRequest `json:"object_attributes"`
}
