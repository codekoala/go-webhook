package gitlab

type WikiPageHook struct {
	Hook

	User             User     `json:"user"`
	Project          Project  `json:"project"`
	Wiki             Wiki     `json:"wiki"`
	ObjectAttributes WikiPage `json:"object_attributes"`
}
