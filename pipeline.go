package gitlab

type Pipeline struct {
	Id         int      `json:"id"`
	Ref        string   `json:"ref"`
	Tag        bool     `json:"tag"`
	SHA        string   `json:"sha"`
	BeforeSHA  string   `json:"before_sha"`
	Status     string   `json:"status"`
	Stages     []string `json:"stages"`
	CreatedAt  string   `json:"created_at"`
	FinishedAt string   `json:"finished_at"`
	Duration   *int     `json:"duration"`
}
