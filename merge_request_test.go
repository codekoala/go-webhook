package gitlab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeRequestHook(t *testing.T) {
	jsonString := `
{
  "object_kind": "merge_request",
  "user": {
    "name": "Administrator",
    "username": "root",
    "avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40\u0026d=identicon"
  },
  "project": {
    "id": 1,
    "name":"Gitlab Test",
    "description":"Aut reprehenderit ut est.",
    "web_url":"http://example.com/gitlabhq/gitlab-test",
    "avatar_url":null,
    "git_ssh_url":"git@example.com:gitlabhq/gitlab-test.git",
    "git_http_url":"http://example.com/gitlabhq/gitlab-test.git",
    "namespace":"GitlabHQ",
    "visibility_level":20,
    "path_with_namespace":"gitlabhq/gitlab-test",
    "default_branch":"master",
    "homepage":"http://example.com/gitlabhq/gitlab-test",
    "url":"http://example.com/gitlabhq/gitlab-test.git",
    "ssh_url":"git@example.com:gitlabhq/gitlab-test.git",
    "http_url":"http://example.com/gitlabhq/gitlab-test.git"
  },
  "repository": {
    "name": "Gitlab Test",
    "url": "http://example.com/gitlabhq/gitlab-test.git",
    "description": "Aut reprehenderit ut est.",
    "homepage": "http://example.com/gitlabhq/gitlab-test"
  },
  "object_attributes": {
    "id": 99,
    "target_branch": "master",
    "source_branch": "ms-viewport",
    "source_project_id": 14,
    "author_id": 51,
    "assignee_id": 6,
    "title": "MS-Viewport",
    "created_at": "2013-12-03T17:23:34Z",
    "updated_at": "2013-12-03T17:23:34Z",
    "milestone_id": null,
    "state": "opened",
    "merge_status": "unchecked",
    "target_project_id": 14,
    "iid": 1,
    "description": "",
    "source": {
      "name":"Awesome Project",
      "description":"Aut reprehenderit ut est.",
      "web_url":"http://example.com/awesome_space/awesome_project",
      "avatar_url":null,
      "git_ssh_url":"git@example.com:awesome_space/awesome_project.git",
      "git_http_url":"http://example.com/awesome_space/awesome_project.git",
      "namespace":"Awesome Space",
      "visibility_level":20,
      "path_with_namespace":"awesome_space/awesome_project",
      "default_branch":"master",
      "homepage":"http://example.com/awesome_space/awesome_project",
      "url":"http://example.com/awesome_space/awesome_project.git",
      "ssh_url":"git@example.com:awesome_space/awesome_project.git",
      "http_url":"http://example.com/awesome_space/awesome_project.git"
    },
    "target": {
      "name":"Awesome Project",
      "description":"Aut reprehenderit ut est.",
      "web_url":"http://example.com/awesome_space/awesome_project",
      "avatar_url":null,
      "git_ssh_url":"git@example.com:awesome_space/awesome_project.git",
      "git_http_url":"http://example.com/awesome_space/awesome_project.git",
      "namespace":"Awesome Space",
      "visibility_level":20,
      "path_with_namespace":"awesome_space/awesome_project",
      "default_branch":"master",
      "homepage":"http://example.com/awesome_space/awesome_project",
      "url":"http://example.com/awesome_space/awesome_project.git",
      "ssh_url":"git@example.com:awesome_space/awesome_project.git",
      "http_url":"http://example.com/awesome_space/awesome_project.git"
    },
    "last_commit": {
      "id": "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
      "message": "fixed readme",
      "timestamp": "2012-01-03T23:36:29+02:00",
      "url": "http://example.com/awesome_space/awesome_project/commits/da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
      "author": {
        "name": "GitLab dev user",
        "email": "gitlabdev@dv6700.(none)"
      }
    },
    "work_in_progress": false,
    "url": "http://example.com/diaspora/merge_requests/1",
    "action": "open",
    "assignee": {
      "name": "User1",
      "username": "user1",
      "avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40\u0026d=identicon"
    }
  },
  "labels": [{
    "id": 206,
    "title": "API",
    "color": "#ffffff",
    "project_id": 14,
    "created_at": "2013-12-03T17:15:43Z",
    "updated_at": "2013-12-03T17:15:43Z",
    "template": false,
    "description": "API related issues",
    "type": "ProjectLabel",
    "group_id": 41
  }],
  "changes": {
    "updated_by_id": [null, 1],
    "updated_at": ["2017-09-15 16:50:55 UTC", "2017-09-15 16:52:00 UTC"],
    "labels": {
      "previous": [{
        "id": 206,
        "title": "API",
        "color": "#ffffff",
        "project_id": 14,
        "created_at": "2013-12-03T17:15:43Z",
        "updated_at": "2013-12-03T17:15:43Z",
        "template": false,
        "description": "API related issues",
        "type": "ProjectLabel",
        "group_id": 41
      }],
      "current": [{
        "id": 205,
        "title": "Platform",
        "color": "#123123",
        "project_id": 14,
        "created_at": "2013-12-03T17:15:43Z",
        "updated_at": "2013-12-03T17:15:43Z",
        "template": false,
        "description": "Platform related issues",
        "type": "ProjectLabel",
        "group_id": 41
      }]
    }
  }
}
`
	kind, hook, err := ParseBody([]byte(jsonString))
	mrHook, ok := hook.(*MergeRequestHook)
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, kind, "merge_request")
	assert.Equal(t, mrHook.ObjectKind, "merge_request")

	assert.Equal(t, mrHook.User.Name, "Administrator")
	assert.Equal(t, mrHook.User.Username, "root")
	assert.Equal(t, mrHook.User.AvatarUrl, "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40\u0026d=identicon")

	assert.Equal(t, mrHook.Project.ID, 1)
	assert.Equal(t, mrHook.Project.Name, "Gitlab Test")
	assert.Equal(t, mrHook.Project.Description, "Aut reprehenderit ut est.")
	assert.Equal(t, mrHook.Project.WebUrl, "http://example.com/gitlabhq/gitlab-test")
	assert.Nil(t, mrHook.Project.AvatarUrl)
	assert.Equal(t, mrHook.Project.GitSshUrl, "git@example.com:gitlabhq/gitlab-test.git")
	assert.Equal(t, mrHook.Project.GitHttpUrl, "http://example.com/gitlabhq/gitlab-test.git")
	assert.Equal(t, mrHook.Project.Namespace, "GitlabHQ")
	assert.Equal(t, mrHook.Project.VisibilityLevel, 20)
	assert.Equal(t, mrHook.Project.PathWithNamespace, "gitlabhq/gitlab-test")
	assert.Equal(t, mrHook.Project.DefaultBranch, "master")
	assert.Equal(t, mrHook.Project.Homepage, "http://example.com/gitlabhq/gitlab-test")
	assert.Equal(t, mrHook.Project.Url, "http://example.com/gitlabhq/gitlab-test.git")
	assert.Equal(t, mrHook.Project.SshUrl, "git@example.com:gitlabhq/gitlab-test.git")
	assert.Equal(t, mrHook.Project.HttpUrl, "http://example.com/gitlabhq/gitlab-test.git")

	assert.Equal(t, mrHook.Repository.Name, "Gitlab Test")
	assert.Equal(t, mrHook.Repository.Url, "http://example.com/gitlabhq/gitlab-test.git")
	assert.Equal(t, mrHook.Repository.Description, "Aut reprehenderit ut est.")
	assert.Equal(t, mrHook.Repository.Homepage, "http://example.com/gitlabhq/gitlab-test")

	mr := mrHook.ObjectAttributes
	assert.Equal(t, mr.Id, 99)
	assert.Equal(t, mr.TargetBranch, "master")
	assert.Equal(t, mr.SourceBranch, "ms-viewport")
	assert.Equal(t, mr.SourceProjectID, 14)
	assert.Equal(t, mr.AuthorID, 51)
	assert.Equal(t, *mr.AssigneeID, 6)
	assert.Equal(t, mr.Title, "MS-Viewport")
	assert.Equal(t, mr.CreatedAt, "2013-12-03T17:23:34Z")
	assert.Equal(t, mr.UpdatedAt, "2013-12-03T17:23:34Z")
	assert.Nil(t, mr.MilestoneId)
	assert.Equal(t, mr.State, "opened")
	assert.Equal(t, mr.MergeStatus, "unchecked")
	assert.Equal(t, mr.TargetProjectID, 14)
	assert.Equal(t, mr.IId, 1)
	assert.Equal(t, mr.Description, "")

	source := mr.Source
	assert.Equal(t, source.Name, "Awesome Project")
	assert.Equal(t, source.Description, "Aut reprehenderit ut est.")
	assert.Equal(t, source.WebUrl, "http://example.com/awesome_space/awesome_project")
	assert.Nil(t, source.AvatarUrl)
	assert.Equal(t, source.GitSshUrl, "git@example.com:awesome_space/awesome_project.git")
	assert.Equal(t, source.GitHttpUrl, "http://example.com/awesome_space/awesome_project.git")
	assert.Equal(t, source.Namespace, "Awesome Space")
	assert.Equal(t, source.VisibilityLevel, 20)
	assert.Equal(t, source.PathWithNamespace, "awesome_space/awesome_project")
	assert.Equal(t, source.DefaultBranch, "master")
	assert.Equal(t, source.Homepage, "http://example.com/awesome_space/awesome_project")
	assert.Equal(t, source.Url, "http://example.com/awesome_space/awesome_project.git")
	assert.Equal(t, source.SshUrl, "git@example.com:awesome_space/awesome_project.git")
	assert.Equal(t, source.HttpUrl, "http://example.com/awesome_space/awesome_project.git")

	target := mr.Target
	assert.Equal(t, target.Name, "Awesome Project")
	assert.Equal(t, target.Description, "Aut reprehenderit ut est.")
	assert.Equal(t, target.WebUrl, "http://example.com/awesome_space/awesome_project")
	assert.Nil(t, target.AvatarUrl)
	assert.Equal(t, target.GitSshUrl, "git@example.com:awesome_space/awesome_project.git")
	assert.Equal(t, target.GitHttpUrl, "http://example.com/awesome_space/awesome_project.git")
	assert.Equal(t, target.Namespace, "Awesome Space")
	assert.Equal(t, target.VisibilityLevel, 20)
	assert.Equal(t, target.PathWithNamespace, "awesome_space/awesome_project")
	assert.Equal(t, target.DefaultBranch, "master")
	assert.Equal(t, target.Homepage, "http://example.com/awesome_space/awesome_project")
	assert.Equal(t, target.Url, "http://example.com/awesome_space/awesome_project.git")
	assert.Equal(t, target.SshUrl, "git@example.com:awesome_space/awesome_project.git")
	assert.Equal(t, target.HttpUrl, "http://example.com/awesome_space/awesome_project.git")

	commit := mr.LastCommit
	assert.Equal(t, commit.Id, "da1560886d4f094c3e6c9ef40349f7d38b5d27d7")
	assert.Equal(t, commit.Message, "fixed readme")
	assert.Equal(t, commit.Timestamp, "2012-01-03T23:36:29+02:00")
	assert.Equal(t, commit.Url, "http://example.com/awesome_space/awesome_project/commits/da1560886d4f094c3e6c9ef40349f7d38b5d27d7")
	assert.Equal(t, commit.Author.Name, "GitLab dev user")
	assert.Equal(t, commit.Author.Email, "gitlabdev@dv6700.(none)")

	assert.False(t, mr.WorkInProgress)
	assert.Equal(t, mr.Url, "http://example.com/diaspora/merge_requests/1")
	assert.Equal(t, mr.Action, "open")
	assert.Equal(t, mr.Assignee.Name, "User1")
	assert.Equal(t, mr.Assignee.Username, "user1")
	assert.Equal(t, mr.Assignee.AvatarUrl, "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40\u0026d=identicon")

	labels := mrHook.Labels
	assert.Equal(t, len(labels), 1)
	label := labels[0]
	assert.Equal(t, label.Id, 206)
	assert.Equal(t, label.Title, "API")
	assert.Equal(t, label.Color, "#ffffff")
	assert.Equal(t, label.ProjectID, 14)
	assert.Equal(t, label.CreatedAt, "2013-12-03T17:15:43Z")
	assert.Equal(t, label.UpdatedAt, "2013-12-03T17:15:43Z")
	assert.Equal(t, label.Template, false)
	assert.Equal(t, label.Description, "API related issues")
	assert.Equal(t, label.Type, "ProjectLabel")
	assert.Equal(t, label.GroupID, 41)

	changes := mrHook.Changes
	assert.Equal(t, len(changes.UpdatedById), 2)
	assert.Nil(t, changes.UpdatedById[0])
	assert.Equal(t, *(changes.UpdatedById[1]), 1)
	assert.Equal(t, len(changes.UpdatedAt), 2)
	assert.Equal(t, changes.UpdatedAt[0], "2017-09-15 16:50:55 UTC")
	assert.Equal(t, changes.UpdatedAt[1], "2017-09-15 16:52:00 UTC")
	assert.Equal(t, len(changes.Labels.Previous), 1)
	assert.Equal(t, changes.Labels.Previous[0].Id, 206)
	assert.Equal(t, changes.Labels.Previous[0].Title, "API")
	assert.Equal(t, changes.Labels.Previous[0].Color, "#ffffff")
	assert.Equal(t, changes.Labels.Previous[0].ProjectID, 14)
	assert.Equal(t, changes.Labels.Previous[0].CreatedAt, "2013-12-03T17:15:43Z")
	assert.Equal(t, changes.Labels.Previous[0].UpdatedAt, "2013-12-03T17:15:43Z")
	assert.False(t, changes.Labels.Previous[0].Template)
	assert.Equal(t, changes.Labels.Previous[0].Description, "API related issues")
	assert.Equal(t, changes.Labels.Previous[0].Type, "ProjectLabel")
	assert.Equal(t, changes.Labels.Previous[0].GroupID, 41)
	assert.Equal(t, len(changes.Labels.Current), 1)
	assert.Equal(t, changes.Labels.Current[0].Id, 205)
	assert.Equal(t, changes.Labels.Current[0].Title, "Platform")
	assert.Equal(t, changes.Labels.Current[0].Color, "#123123")
	assert.Equal(t, changes.Labels.Current[0].ProjectID, 14)
	assert.Equal(t, changes.Labels.Current[0].CreatedAt, "2013-12-03T17:15:43Z")
	assert.Equal(t, changes.Labels.Current[0].UpdatedAt, "2013-12-03T17:15:43Z")
	assert.False(t, changes.Labels.Current[0].Template)
	assert.Equal(t, changes.Labels.Current[0].Description, "Platform related issues")
	assert.Equal(t, changes.Labels.Current[0].Type, "ProjectLabel")
	assert.Equal(t, changes.Labels.Current[0].GroupID, 41)
}
