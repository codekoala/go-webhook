package gitlab

type Wiki struct {
	WebUrl            string `json:"web_url"`
	GitSshUrl         string `json:"git_ssh_url"`
	GitHttpUrl        string `json:"git_http_url"`
	PathWithNamespace string `json:"path_with_namespace"`
	DefaultBranch     string `json:"default_branch"`
}

type WikiPage struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Format  string `json:"format"`
	Message string `json:"message"`
	Slug    string `json:"slug"`
	Url     string `json:"url"`
	Action  string `json:"action"`
}
