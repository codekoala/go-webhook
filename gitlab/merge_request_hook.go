package gitlab

type MergeRequestHook struct {
	Hook

	User             User         `json:"user"`
	ObjectAttributes MergeRequest `json:"object_attributes"`
}
