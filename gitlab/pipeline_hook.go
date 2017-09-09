package gitlab

type PipelineHook struct {
	Hook

	ObjectAttributes Pipeline `json:"object_attributes"`
	User             User     `json:"user"`
	Project          Project  `json:"project"`
	Commit           Commit   `json:"commit"`
	Builds           []Build  `json:"builds"`
}
