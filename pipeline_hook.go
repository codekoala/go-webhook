package gitlab

// See https://gitlab.com/help/user/project/integrations/webhooks.md#pipeline-events
type PipelineHook struct {
	Hook

	ObjectAttributes Pipeline `json:"object_attributes"`
	User             User     `json:"user"`
	Project          Project  `json:"project"`
	Commit           Commit   `json:"commit"`
	Builds           []Build  `json:"builds"`
}
