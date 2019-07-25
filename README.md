# go-webhook
Helpers to deal with GitLab webhooks in Go.

Adapted GitLab Version: `CE 11.9.8`

## Example Usage

```go
// request send by GitLab
var body []byte

kind, hook, err := ParseBody(body)

if mrHook, ok := hook.(*MergeRequestHook); ok {
	// do something
}
```

## Event Adapted to Current GitLab Version

- [x] Push
- [x] Merge Request
