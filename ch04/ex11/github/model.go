package github

// POST /repos/:owner/:repo/issues 対応struct
// Assignee, Assignees, Milestone は除外する
type CreateIssueParameter struct {
	Title string
	Body string
	Labels []string
}
// PATCH /repos/:owner/:repo/issues/:issue_number 対応struct
// Assignee, Assignees, Milestone は除外する
type EditIssueParameter struct {
	Title string
	Body string
	State string
	Labels []string
}
