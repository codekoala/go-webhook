# gitlab-webhook-go
Helpers to deal with GitLab webhooks in Go.

Adapted GitLab Version: `CE 11.9.8`

## Example Usage

```go
import (
	"github.com/vimsucks/gitlab-webhook-go/gitlab"
)
// request send by GitLab

func foo(body []byte) {
    kind, hook, err := gitlab.ParseBody(body)

    if mrHook, ok := hook.(*MergeRequestHook); ok {
        // do something
    }
}
```

## Unit Test Covered Events

- [x] Push `PushHook`
- [x] Tag `TagHook`
- [x] Issue `IssueHook`
- [x] Comment on Commit `NoteHook`
- [ ] Comment on Merge Request
- [ ] Comment on Issue
- [ ] Comment on Code Snippet
- [x] Merge Request `MergeRequestHook`
- [ ] Wiki Page
- [ ] Pipeline
- [ ] Build
