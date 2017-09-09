package gitlab

type Project struct {
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	WebUrl            string  `json:"web_url"`
	AvatarUrl         *string `json:"avatar_url"`
	GitSshUrl         string  `json:"git_ssh_url"`
	GitHttpUrl        string  `json:"git_http_url"`
	Namespace         string  `json:"namespace"`
	VisibilityLevel   int     `json:"visibility_level"`
	PathWithNamespace string  `json:"path_with_namespace"`
	DefaultBranch     string  `json:"default_branch"`
	Homepage          string  `json:"homepage"`
	Url               string  `json:"url"`
	SshUrl            string  `json:"ssh_url"`
	HttpUrl           string  `json:"http_url"`
}
