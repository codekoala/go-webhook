package gitlab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommitNote(t *testing.T) {
	json := `
{
  "object_kind": "note",
  "user": {
    "name": "Administrator",
    "username": "root",
    "avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40\u0026d=identicon"
  },
  "project_id": 5,
  "project":{
    "id": 5,
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
  "repository":{
    "name": "Gitlab Test",
    "url": "http://example.com/gitlab-org/gitlab-test.git",
    "description": "Aut reprehenderit ut est.",
    "homepage": "http://example.com/gitlab-org/gitlab-test"
  },
  "object_attributes": {
    "id": 1243,
    "note": "This is a commit comment. How does this work?",
    "noteable_type": "Commit",
    "author_id": 1,
    "created_at": "2015-05-17 18:08:09 UTC",
    "updated_at": "2015-05-17 18:08:09 UTC",
    "project_id": 5,
    "attachment":null,
    "line_code": "bec9703f7a456cd2b4ab5fb3220ae016e3e394e3_0_1",
    "commit_id": "cfe32cf61b73a0d5e9f13e774abde7ff789b1660",
    "noteable_id": null,
    "system": false,
    "st_diff": {
      "diff": "--- /dev/null\n+++ b/six\n@@ -0,0 +1 @@\n+Subproject commit 409f37c4f05865e4fb208c771485f211a22c4c2d\n",
      "new_path": "six",
      "old_path": "six",
      "a_mode": "0",
      "b_mode": "160000",
      "new_file": true,
      "renamed_file": false,
      "deleted_file": false
    },
    "url": "http://example.com/gitlab-org/gitlab-test/commit/cfe32cf61b73a0d5e9f13e774abde7ff789b1660#note_1243"
  },
  "commit": {
    "id": "cfe32cf61b73a0d5e9f13e774abde7ff789b1660",
    "message": "Add submodule\n\nSigned-off-by: Dmitriy Zaporozhets \u003cdmitriy.zaporozhets@gmail.com\u003e\n",
    "timestamp": "2014-02-27T10:06:20+02:00",
    "url": "http://example.com/gitlab-org/gitlab-test/commit/cfe32cf61b73a0d5e9f13e774abde7ff789b1660",
    "author": {
      "name": "Dmitriy Zaporozhets",
      "email": "dmitriy.zaporozhets@gmail.com"
    }
  }
}
`
	kind, hook, err := ParseBody([]byte(json))
	issueNoteHook, ok := hook.(*NoteHook)
	assert.Nil(t, err)
	assert.NotNil(t, issueNoteHook)
	assert.True(t, ok)
	assert.Equal(t, kind, "note")
	user := User{
		Name:      "Administrator",
		Username:  "root",
		Email:     "",
		AvatarUrl: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40\u0026d=identicon",
	}
	assert.Equal(t, user, issueNoteHook.User)
	assert.Equal(t, issueNoteHook.ProjectID, 5)
	project := Project{
		ID:                5,
		Name:              "Gitlab Test",
		Description:       "Aut reprehenderit ut est.",
		WebUrl:            "http://example.com/gitlabhq/gitlab-test",
		AvatarUrl:         nil,
		GitSshUrl:         "git@example.com:gitlabhq/gitlab-test.git",
		GitHttpUrl:        "http://example.com/gitlabhq/gitlab-test.git",
		Namespace:         "GitlabHQ",
		VisibilityLevel:   20,
		PathWithNamespace: "gitlabhq/gitlab-test",
		DefaultBranch:     "master",
		Homepage:          "http://example.com/gitlabhq/gitlab-test",
		Url:               "http://example.com/gitlabhq/gitlab-test.git",
		SshUrl:            "git@example.com:gitlabhq/gitlab-test.git",
		HttpUrl:           "http://example.com/gitlabhq/gitlab-test.git",
	}
	assert.Equal(t, issueNoteHook.Project, project)
	repository := Repository{
		Name:            "Gitlab Test",
		Url:             "http://example.com/gitlab-org/gitlab-test.git",
		Description:     "Aut reprehenderit ut est.",
		Homepage:        "http://example.com/gitlab-org/gitlab-test",
		GitSshUrl:       "",
		GitHttpUrl:      "",
		VisibilityLevel: 0,
	}
	assert.Equal(t, issueNoteHook.Repository, repository)
	lineCode := "bec9703f7a456cd2b4ab5fb3220ae016e3e394e3_0_1"
	comment := Note{
		Id:           1243,
		Note:         "This is a commit comment. How does this work?",
		NoteableType: "Commit",
		AuthorID:     1,
		CreatedAt:    "2015-05-17 18:08:09 UTC",
		UpdatedAt:    "2015-05-17 18:08:09 UTC",
		ProjectID:    5,
		Attachment:   nil,
		LineCode:     &lineCode,
		CommitID:     "cfe32cf61b73a0d5e9f13e774abde7ff789b1660",
		NoteableID:   nil,
		System:       false,
		StDiff: &Diff{
			Diff:        "--- /dev/null\n+++ b/six\n@@ -0,0 +1 @@\n+Subproject commit 409f37c4f05865e4fb208c771485f211a22c4c2d\n",
			NewPath:     "six",
			OldPath:     "six",
			AMode:       "0",
			BMode:       "160000",
			NewFile:     true,
			RenamedFile: false,
			DeletedFile: false,
		},
		Url: "http://example.com/gitlab-org/gitlab-test/commit/cfe32cf61b73a0d5e9f13e774abde7ff789b1660#note_1243",
	}
	assert.Equal(t, issueNoteHook.ObjectAttributes, comment)
	commit := Commit{
		Id:        "cfe32cf61b73a0d5e9f13e774abde7ff789b1660",
		Message:   "Add submodule\n\nSigned-off-by: Dmitriy Zaporozhets \u003cdmitriy.zaporozhets@gmail.com\u003e\n",
		Timestamp: "2014-02-27T10:06:20+02:00",
		Url:       "http://example.com/gitlab-org/gitlab-test/commit/cfe32cf61b73a0d5e9f13e774abde7ff789b1660",
		Author: User{
			Name:  "Dmitriy Zaporozhets",
			Email: "dmitriy.zaporozhets@gmail.com",
		},
		Added:    nil,
		Modified: nil,
		Removed:  nil,
	}
	assert.Equal(t, issueNoteHook.Commit, &commit)
	assert.Nil(t, issueNoteHook.Issue)
	assert.Nil(t, issueNoteHook.MergeRequest)
	assert.Nil(t, issueNoteHook.Snippet)
}
