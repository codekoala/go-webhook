package gitlab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTagHook(t *testing.T) {
	jsonString := `
{
  "object_kind": "tag_push",
  "before": "0000000000000000000000000000000000000000",
  "after": "82b3d5ae55f7080f1e6022629cdb57bfae7cccc7",
  "ref": "refs/tags/v1.0.0",
  "checkout_sha": "82b3d5ae55f7080f1e6022629cdb57bfae7cccc7",
  "user_id": 1,
  "user_name": "John Smith",
  "user_avatar": "https://s.gravatar.com/avatar/d4c74594d841139328695756648b6bd6?s=8://s.gravatar.com/avatar/d4c74594d841139328695756648b6bd6?s=80",
  "project_id": 1,
  "project":{
    "id": 1,
    "name":"Example",
    "description":"",
    "web_url":"http://example.com/jsmith/example",
    "avatar_url":null,
    "git_ssh_url":"git@example.com:jsmith/example.git",
    "git_http_url":"http://example.com/jsmith/example.git",
    "namespace":"Jsmith",
    "visibility_level":0,
    "path_with_namespace":"jsmith/example",
    "default_branch":"master",
    "homepage":"http://example.com/jsmith/example",
    "url":"git@example.com:jsmith/example.git",
    "ssh_url":"git@example.com:jsmith/example.git",
    "http_url":"http://example.com/jsmith/example.git"
  },
  "repository":{
    "name": "Example",
    "url": "ssh://git@example.com/jsmith/example.git",
    "description": "",
    "homepage": "http://example.com/jsmith/example",
    "git_http_url":"http://example.com/jsmith/example.git",
    "git_ssh_url":"git@example.com:jsmith/example.git",
    "visibility_level":0
  },
  "commits": [],
  "total_commits_count": 0
}
`
	kind, hook, err := ParseBody([]byte(jsonString))
	tag, ok := hook.(*TagHook)
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, kind, "tag_push")
	assert.Equal(t, tag.ObjectKind, "tag_push")
	assert.Equal(t, tag.Before, "0000000000000000000000000000000000000000")
	assert.Equal(t, tag.After, "82b3d5ae55f7080f1e6022629cdb57bfae7cccc7")
	assert.Equal(t, tag.Ref, "refs/tags/v1.0.0")
	assert.Equal(t, tag.CheckoutSHA, "82b3d5ae55f7080f1e6022629cdb57bfae7cccc7")
	assert.Equal(t, tag.UserID, 1)
	assert.Equal(t, tag.UserName, "John Smith")
	assert.Equal(t, tag.UserAvatar, "https://s.gravatar.com/avatar/d4c74594d841139328695756648b6bd6?s=8://s.gravatar.com/avatar/d4c74594d841139328695756648b6bd6?s=80")
	assert.Equal(t, tag.ProjectID, 1)

	assert.Equal(t, tag.Project.ID, 1)
	assert.Equal(t, tag.Project.Name, "Example")
	assert.Equal(t, tag.Project.Description, "")
	assert.Equal(t, tag.Project.WebUrl, "http://example.com/jsmith/example")
	assert.Nil(t, tag.Project.AvatarUrl)
	assert.Equal(t, tag.Project.GitSshUrl, "git@example.com:jsmith/example.git")
	assert.Equal(t, tag.Project.GitHttpUrl, "http://example.com/jsmith/example.git")
	assert.Equal(t, tag.Project.Namespace, "Jsmith")
	assert.Equal(t, tag.Project.VisibilityLevel, 0)
	assert.Equal(t, tag.Project.PathWithNamespace, "jsmith/example")
	assert.Equal(t, tag.Project.DefaultBranch, "master")
	assert.Equal(t, tag.Project.Homepage, "http://example.com/jsmith/example")
	assert.Equal(t, tag.Project.Url, "git@example.com:jsmith/example.git")
	assert.Equal(t, tag.Project.SshUrl, "git@example.com:jsmith/example.git")
	assert.Equal(t, tag.Project.HttpUrl, "http://example.com/jsmith/example.git")

	assert.Equal(t, tag.Repository.Name, "Example")
	assert.Equal(t, tag.Repository.Url, "ssh://git@example.com/jsmith/example.git")
	assert.Equal(t, tag.Repository.Description, "")
	assert.Equal(t, tag.Repository.Homepage, "http://example.com/jsmith/example")
	assert.Equal(t, tag.Repository.GitHttpUrl, "http://example.com/jsmith/example.git")
	assert.Equal(t, tag.Repository.GitSshUrl, "git@example.com:jsmith/example.git")
	assert.Equal(t, tag.Repository.VisibilityLevel, 0)

	commits := tag.Commits
	assert.Equal(t, len(commits), 0)
	assert.Equal(t, tag.TotalCommitsCount, 0)
}
