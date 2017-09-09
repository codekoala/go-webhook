# go-webhook
Helpers to deal with webhooks in Go.

The primary purpose of this repository is to create Go structs to hold data from webhook request payloads. 
Currently only support for [GitLab](https://docs.gitlab.com/ce/user/project/integrations/webhooks.html) has 
been implemented.

## Example Usage

See the `main.go` in the repository root for an example for how to work with GitLab webhook requests.
