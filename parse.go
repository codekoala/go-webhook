package gitlab

import (
	"io"
	"io/ioutil"

	"github.com/json-iterator/go"
)

// ParseBody parses a GitLab hook from the specified JSON document
func ParseBody(body []byte) (kind string, hook interface{}, err error) {
	switch kind = jsoniter.Get(body, "object_kind").ToString(); kind {
	case "push":
		hook = new(PushHook)
	case "tag_push":
		hook = new(TagHook)
	case "issue":
		hook = new(IssueHook)
	case "note":
		hook = new(NoteHook)
	case "merge_request":
		hook = new(MergeRequestHook)
	case "wiki_page":
		hook = new(WikiPageHook)
	case "pipeline":
		hook = new(PipelineHook)
	case "build":
		hook = new(BuildHook)
	}

	if err = jsoniter.Unmarshal(body, hook); err != nil {
		return
	}

	return kind, hook, nil
}

// ParseReader reads everything from the specified io.Reader and attempts to parse it as a JSON document with a
// GitLab hook payload.
func ParseReader(in io.Reader) (kind string, hook interface{}, err error) {
	var body []byte

	if body, err = ioutil.ReadAll(in); err != nil {
		return
	}

	return ParseBody(body)
}
