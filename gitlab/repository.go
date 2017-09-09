package gitlab

type Repository struct {
	Name            string `json:"name"`
	Url             string `json:"url"`
	Description     string `json:"description"`
	Homepage        string `json:"homepage"`
	GitSshUrl       string `json:"git_ssh_url"`
	GitHttpUrl      string `json:"git_http_url"`
	VisibilityLevel int    `json:"visibility_level"`
}
