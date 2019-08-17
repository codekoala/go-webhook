package gitlab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIssueHook(t *testing.T) {
	jsonString := `{
  "object_kind": "issue",
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
    "id": 301,
    "title": "New API: create/update/delete file",
    "assignee_ids": [51],
    "assignee_id": 51,
    "author_id": 51,
    "project_id": 14,
    "created_at": "2013-12-03T17:15:43Z",
    "updated_at": "2013-12-03T17:15:43Z",
    "position": 0,
    "branch_name": null,
    "description": "Create new API for manipulations with repository",
    "milestone_id": null,
    "state": "opened",
    "iid": 23,
    "url": "http://example.com/diaspora/issues/23",
    "action": "open"
  },
  "assignees": [{
    "name": "User1",
    "username": "user1",
    "avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40\u0026d=identicon"
  }],
  "assignee": {
    "name": "User1",
    "username": "user1",
    "avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40\u0026d=identicon"
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
}`
	kind, hook, err := ParseBody([]byte(jsonString))
	issueHook, ok := hook.(*IssueHook)
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, kind, "issue")
	assert.Equal(t, issueHook.ObjectKind, "issue")

	user := issueHook.User
	assert.Equal(t, user.Name, "Administrator")
	assert.Equal(t, user.Username, "root")
	assert.Equal(t, user.AvatarUrl, "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40\u0026d=identicon")

	project := issueHook.Project
	assert.Equal(t, project.ID, 1)
	assert.Equal(t, project.Name, "Gitlab Test")
	assert.Equal(t, project.Description, "Aut reprehenderit ut est.")
	assert.Equal(t, project.WebUrl, "http://example.com/gitlabhq/gitlab-test")
	assert.Nil(t, project.AvatarUrl)
	assert.Equal(t, project.GitSshUrl, "git@example.com:gitlabhq/gitlab-test.git")
	assert.Equal(t, project.GitHttpUrl, "http://example.com/gitlabhq/gitlab-test.git")
	assert.Equal(t, project.Namespace, "GitlabHQ")
	assert.Equal(t, project.VisibilityLevel, 20)
	assert.Equal(t, project.PathWithNamespace, "gitlabhq/gitlab-test")
	assert.Equal(t, project.DefaultBranch, "master")
	assert.Equal(t, project.Homepage, "http://example.com/gitlabhq/gitlab-test")
	assert.Equal(t, project.Url, "http://example.com/gitlabhq/gitlab-test.git")
	assert.Equal(t, project.SshUrl, "git@example.com:gitlabhq/gitlab-test.git")
	assert.Equal(t, project.HttpUrl, "http://example.com/gitlabhq/gitlab-test.git")

	repository := issueHook.Repository
	assert.Equal(t, repository.Name, "Gitlab Test")
	assert.Equal(t, repository.Url, "http://example.com/gitlabhq/gitlab-test.git")
	assert.Equal(t, repository.Description, "Aut reprehenderit ut est.")
	assert.Equal(t, repository.Homepage, "http://example.com/gitlabhq/gitlab-test")

	objectAttributes := issueHook.ObjectAttributes
	assert.Equal(t, objectAttributes.Id, 301)
	assert.Equal(t, objectAttributes.Title, "New API: create/update/delete file")
	assert.Equal(t, objectAttributes.AssigneeIDs, []int{51})
	assert.Equal(t, *objectAttributes.AssigneeID, 51)
	assert.Equal(t, objectAttributes.AuthorID, 51)
	assert.Equal(t, objectAttributes.ProjectID, 14)
	assert.Equal(t, objectAttributes.CreatedAt, "2013-12-03T17:15:43Z")
	assert.Equal(t, objectAttributes.UpdatedAt, "2013-12-03T17:15:43Z")
	assert.Equal(t, objectAttributes.Position, 0)
	assert.Nil(t, objectAttributes.BranchName)
	assert.Equal(t, objectAttributes.Description, "Create new API for manipulations with repository")
	assert.Nil(t, objectAttributes.MilestoneId)
	assert.Equal(t, objectAttributes.State, "opened")
	assert.Equal(t, objectAttributes.IId, 23)
	assert.Equal(t, objectAttributes.Url, "http://example.com/diaspora/issues/23")
	assert.Equal(t, objectAttributes.Action, "open")

	assignee := issueHook.Assignee
	assert.Equal(t, assignee.Name, "User1")
	assert.Equal(t, assignee.Username, "user1")
	assert.Equal(t, assignee.AvatarUrl, "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40\u0026d=identicon")

	assert.Equal(t, issueHook.Assignees, []User{*assignee})

	label := Label{
		Id:          206,
		Title:       "API",
		Color:       "#ffffff",
		ProjectID:   14,
		CreatedAt:   "2013-12-03T17:15:43Z",
		UpdatedAt:   "2013-12-03T17:15:43Z",
		Template:    false,
		Description: "API related issues",
		Type:        "ProjectLabel",
		GroupID:     41,
	}
	assert.Equal(t, issueHook.Labels, []Label{label})

	changes := issueHook.Changes
	assert.Equal(t, len(changes.UpdatedById), 2)
	assert.Nil(t, changes.UpdatedById[0])
	assert.Equal(t, *changes.UpdatedById[1], 1)
	assert.Equal(t, changes.UpdatedAt, []string{"2017-09-15 16:50:55 UTC", "2017-09-15 16:52:00 UTC"})
	changeLables := changes.Labels
	previous := Label{
		Id:          206,
		Title:       "API",
		Color:       "#ffffff",
		ProjectID:   14,
		CreatedAt:   "2013-12-03T17:15:43Z",
		UpdatedAt:   "2013-12-03T17:15:43Z",
		Template:    false,
		Description: "API related issues",
		Type:        "ProjectLabel",
		GroupID:     41,
	}
	assert.Equal(t, changeLables.Previous, []Label{previous})
	current := Label{
		Id:          205,
		Title:       "Platform",
		Color:       "#123123",
		ProjectID:   14,
		CreatedAt:   "2013-12-03T17:15:43Z",
		UpdatedAt:   "2013-12-03T17:15:43Z",
		Template:    false,
		Description: "Platform related issues",
		Type:        "ProjectLabel",
		GroupID:     41,
	}
	assert.Equal(t, changeLables.Current, []Label{current})

}
