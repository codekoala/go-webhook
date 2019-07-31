# go-webhook
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

## Event Adapted to Current GitLab Version(unit test covered)

- [x] Push
- [x] Merge Request
- [x] Tag
