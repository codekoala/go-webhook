package gitlab

type ArtifactFile struct {
	Filename *string `json:"filename"`
	Size     *int    `json:"size"`
}
