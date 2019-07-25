package gitlab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPushHook(t *testing.T) {
	jsonString := `
{
  "object_kind": "push",
  "before": "95790bf891e76fee5e1747ab589903a6a1f80f22",
  "after": "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
  "ref": "refs/heads/master",
  "checkout_sha": "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
  "user_id": 4,
  "user_name": "John Smith",
  "user_username": "jsmith",
  "user_email": "john@example.com",
  "user_avatar": "https://s.gravatar.com/avatar/d4c74594d841139328695756648b6bd6?s=8://s.gravatar.com/avatar/d4c74594d841139328695756648b6bd6?s=80",
  "project_id": 15,
  "project":{
    "id": 15,
    "name":"Diaspora",
    "description":"",
    "web_url":"http://example.com/mike/diaspora",
    "avatar_url":null,
    "git_ssh_url":"git@example.com:mike/diaspora.git",
    "git_http_url":"http://example.com/mike/diaspora.git",
    "namespace":"Mike",
    "visibility_level":0,
    "path_with_namespace":"mike/diaspora",
    "default_branch":"master",
    "homepage":"http://example.com/mike/diaspora",
    "url":"git@example.com:mike/diaspora.git",
    "ssh_url":"git@example.com:mike/diaspora.git",
    "http_url":"http://example.com/mike/diaspora.git"
  },
  "repository":{
    "name": "Diaspora",
    "url": "git@example.com:mike/diaspora.git",
    "description": "",
    "homepage": "http://example.com/mike/diaspora",
    "git_http_url":"http://example.com/mike/diaspora.git",
    "git_ssh_url":"git@example.com:mike/diaspora.git",
    "visibility_level":0
  },
  "commits": [
    {
      "id": "b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327",
      "message": "Update Catalan translation to e38cb41.",
      "timestamp": "2011-12-12T14:27:31+02:00",
      "url": "http://example.com/mike/diaspora/commit/b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327",
      "author": {
        "name": "Jordi Mallach",
        "email": "jordi@softcatala.org"
      },
      "added": ["CHANGELOG"],
      "modified": ["app/controller/application.rb"],
      "removed": []
    },
    {
      "id": "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
      "message": "fixed readme",
      "timestamp": "2012-01-03T23:36:29+02:00",
      "url": "http://example.com/mike/diaspora/commit/da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
      "author": {
        "name": "GitLab dev user",
        "email": "gitlabdev@dv6700.(none)"
      },
      "added": ["CHANGELOG"],
      "modified": ["app/controller/application.rb"],
      "removed": []
    }
  ],
  "total_commits_count": 4
}
`
	kind, hook, err := ParseBody([]byte(jsonString))
	push, ok := hook.(*PushHook)
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, kind, "push")
	assert.Equal(t, push.ObjectKind, "push")
	assert.Equal(t, push.Before, "95790bf891e76fee5e1747ab589903a6a1f80f22")
	assert.Equal(t, push.After, "da1560886d4f094c3e6c9ef40349f7d38b5d27d7")
	assert.Equal(t, push.Ref, "refs/heads/master")
	assert.Equal(t, push.CheckoutSHA, "da1560886d4f094c3e6c9ef40349f7d38b5d27d7")
	assert.Equal(t, push.UserID, 4)
	assert.Equal(t, push.UserName, "John Smith")
	assert.Equal(t, push.UserUsername, "jsmith")
	assert.Equal(t, push.UserEmail, "john@example.com")
	assert.Equal(t, push.UserAvatar, "https://s.gravatar.com/avatar/d4c74594d841139328695756648b6bd6?s=8://s.gravatar.com/avatar/d4c74594d841139328695756648b6bd6?s=80")
	assert.Equal(t, push.ProjectID, 15)

	assert.Equal(t, push.Project.ID, 15)
	assert.Equal(t, push.Project.Name, "Diaspora")
	assert.Equal(t, push.Project.Description, "")
	assert.Equal(t, push.Project.WebUrl, "http://example.com/mike/diaspora")
	assert.Nil(t, push.Project.AvatarUrl)
	assert.Equal(t, push.Project.GitSshUrl, "git@example.com:mike/diaspora.git")
	assert.Equal(t, push.Project.GitHttpUrl, "http://example.com/mike/diaspora.git")
	assert.Equal(t, push.Project.Namespace, "Mike")
	assert.Equal(t, push.Project.VisibilityLevel, 0)
	assert.Equal(t, push.Project.PathWithNamespace, "mike/diaspora")
	assert.Equal(t, push.Project.DefaultBranch, "master")
	assert.Equal(t, push.Project.Homepage, "http://example.com/mike/diaspora")
	assert.Equal(t, push.Project.Url, "git@example.com:mike/diaspora.git")
	assert.Equal(t, push.Project.SshUrl, "git@example.com:mike/diaspora.git")
	assert.Equal(t, push.Project.HttpUrl, "http://example.com/mike/diaspora.git")

	assert.Equal(t, push.Repository.Name, "Diaspora")
	assert.Equal(t, push.Repository.Url, "git@example.com:mike/diaspora.git")
	assert.Equal(t, push.Repository.Description, "")
	assert.Equal(t, push.Repository.Homepage, "http://example.com/mike/diaspora")
	assert.Equal(t, push.Repository.GitHttpUrl, "http://example.com/mike/diaspora.git")
	assert.Equal(t, push.Repository.GitSshUrl, "git@example.com:mike/diaspora.git")
	assert.Equal(t, push.Repository.VisibilityLevel, 0)

	commits := push.Commits
	assert.Equal(t, len(commits), 2)
	assert.Equal(t, commits[0].Id, "b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327")
	assert.Equal(t, commits[0].Message, "Update Catalan translation to e38cb41.")
	assert.Equal(t, commits[0].Timestamp, "2011-12-12T14:27:31+02:00")
	assert.Equal(t, commits[0].Url, "http://example.com/mike/diaspora/commit/b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327")
	assert.Equal(t, commits[0].Author.Name, "Jordi Mallach")
	assert.Equal(t, commits[0].Author.Email, "jordi@softcatala.org")
	assert.Equal(t, len(commits[1].Added), 1)
	assert.Equal(t, commits[0].Added[0], "CHANGELOG")
	assert.Equal(t, len(commits[0].Modified), 1)
	assert.Equal(t, commits[0].Modified[0], "app/controller/application.rb")
	assert.Equal(t, len(commits[0].Removed), 0)
	assert.Equal(t, commits[1].Id, "da1560886d4f094c3e6c9ef40349f7d38b5d27d7")
	assert.Equal(t, commits[1].Message, "fixed readme")
	assert.Equal(t, commits[1].Timestamp, "2012-01-03T23:36:29+02:00")
	assert.Equal(t, commits[1].Url, "http://example.com/mike/diaspora/commit/da1560886d4f094c3e6c9ef40349f7d38b5d27d7")
	assert.Equal(t, commits[1].Author.Name, "GitLab dev user")
	assert.Equal(t, commits[1].Author.Email, "gitlabdev@dv6700.(none)")
	assert.Equal(t, len(commits[1].Added), 1)
	assert.Equal(t, commits[1].Added[0], "CHANGELOG")
	assert.Equal(t, len(commits[1].Modified), 1)
	assert.Equal(t, commits[1].Modified[0], "app/controller/application.rb")
	assert.Equal(t, len(commits[1].Removed), 0)
	assert.Equal(t, push.TotalCommitsCount, 4)
}
