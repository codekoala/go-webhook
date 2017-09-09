package gitlab

type BuildHook struct {
	Hook

	Ref               string      `json:"ref"`
	Tag               bool        `json:"tag"`
	BeforeSHA         string      `json:"before_sha"`
	SHA               string      `json:"sha"`
	BuildID           int         `json:"build_id"`
	BuildName         string      `json:"build_name"`
	BuildStage        string      `json:"build_stage"`
	BuildStatus       string      `json:"build_status"`
	BuildStartedAt    *string     `json:"build_started_at"`
	BuildFinishedAt   *string     `json:"build_finished_at"`
	BuildDuration     *int        `json:"build_duration"`
	BuildAllowFailure bool        `json:"build_allow_failure"`
	ProjectID         int         `json:"project_id"`
	ProjectName       string      `json:"project_name"`
	User              User        `json:"user"`
	Commit            BuildCommit `json:"commit"`
	Repository        Repository  `json:"repository"`
}
