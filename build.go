package gitlab

type Build struct {
	Id            int          `json:"id"`
	Stage         string       `json:"stage"`
	Name          string       `json:"name"`
	Status        string       `json:"status"`
	CreatedAt     string       `json:"created_at"`
	StartedAt     *string      `json:"started_at"`
	FinishedAt    *string      `json:"finished_at"`
	When          string       `json:"when"`
	Manual        bool         `json:"manual"`
	User          User         `json:"user"`
	Runner        *string      `json:"runner"`
	ArtifactsFile ArtifactFile `json:"artifacts_file"`
}
